package auth

// https://docs.aws.amazon.com/cli/latest/reference/sts/get-session-token.html

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"
	"strings"
)

func GetCredentialsWithMFA(cfg aws.Config, token string, duration int32) (*types.Credentials, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	return GetCredentialsWithMFAWithContext(ctx, cfg, token, duration)
}

func GetCredentialsWithMFAWithContext(ctx context.Context, cfg aws.Config, token string, duration int32) (*types.Credentials, error) {

	sts_opts := sts.Options{
		Region: cfg.Region,
	}

	sts_client := sts.New(sts_opts)

	username, err := username(ctx, sts_client)

	if err != nil {
		return nil, err
	}

	iam_opts := iam.Options{
		Region: cfg.Region,
	}

	iam_client := iam.New(iam_opts)

	mfaDevice, err := mfaDevice(ctx, iam_client, username)

	if err != nil {
		return nil, err
	}

	return sessionCredentials(ctx, sts_client, mfaDevice, token, duration)
}

func username(ctx context.Context, sts_client *sts.Client) (string, error) {

	opts := &sts.GetCallerIdentityInput{}

	rsp, err := sts_client.GetCallerIdentity(ctx, opts)

	if err != nil {
		return "", fmt.Errorf("Failed to derive username, %w", err)
	}

	arn := rsp.Arn

	return strings.Split(*arn, ":user/")[1], nil
}

func mfaDevice(ctx context.Context, iam_client *iam.Client, userArn string) (string, error) {

	opts := &iam.ListMFADevicesInput{
		UserName: &userArn,
	}

	rsp, err := iam_client.ListMFADevices(ctx, opts)

	if err != nil {
		return "", fmt.Errorf("Failed to list devices for %s, %w", userArn, err)
	}

	return *rsp.MFADevices[0].SerialNumber, nil
}

func sessionCredentials(ctx context.Context, sts_client *sts.Client, mfaDevice string, tokenCode string, duration int32) (*types.Credentials, error) {

	opts := &sts.GetSessionTokenInput{
		SerialNumber:    &mfaDevice,
		DurationSeconds: &duration,
		TokenCode:       &tokenCode,
	}

	rsp, err := sts_client.GetSessionToken(ctx, opts)

	if err != nil {
		return nil, fmt.Errorf("Failed to get session token, %w", err)
	}

	creds := rsp.Credentials
	return creds, nil
}
