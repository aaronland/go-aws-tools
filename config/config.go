package config

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	aws_config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"
	"github.com/go-ini/ini"
	"log"
	"os"
)

type Config struct {
	Path string
	ini  *ini.File
}

func NewConfig() (*Config, error) {

	var ini_config *ini.File
	var ini_path string

	for _, path := range aws_config.DefaultSharedCredentialsFiles {

		_, err := os.Stat(path)

		if err != nil {
			continue
		}

		i, err := ini.Load(path)

		if err != nil {
			return nil, err
		}

		ini_path = path
		ini_config = i
		break
	}

	if ini_config == nil {
		return nil, errors.New("Unable to load config file")
	}

	c := Config{
		Path: ini_path,
		ini:  ini_config,
	}

	return &c, nil
}

func (c *Config) AWSConfigWithProfile(profile string) (aws.Config, error) {

	ctx := context.Background()

	sect := c.ini.Section("default")
	region := sect.Key("region")

	log.Println("WHAT", region)

	return aws_config.LoadDefaultConfig(ctx,
		aws_config.WithSharedConfigProfile(profile),
		aws_config.WithRegion(region.String()),
	)
}

func (c *Config) SetSessionCredentialsWithProfile(profile string, creds *types.Credentials) error {

	ini_path := c.Path
	ini_config, err := ini.Load(ini_path)

	if err != nil {
		return err
	}

	sect := ini_config.Section(profile)

	sect.Key("aws_access_key_id").SetValue(*creds.AccessKeyId)
	sect.Key("aws_secret_access_key").SetValue(*creds.SecretAccessKey)
	sect.Key("aws_session_token").SetValue(*creds.SessionToken)

	err = ini_config.SaveTo(ini_path)

	if err != nil {
		return err
	}

	return nil
}
