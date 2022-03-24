// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of aliases (also called CNAMEs or alternate domain names) that
// conflict or overlap with the provided alias, and the associated CloudFront
// distributions and Amazon Web Services accounts for each conflicting alias. In
// the returned list, the distribution and account IDs are partially hidden, which
// allows you to identify the distributions and accounts that you own, but helps to
// protect the information of ones that you don’t own. Use this operation to find
// aliases that are in use in CloudFront that conflict or overlap with the provided
// alias. For example, if you provide www.example.com as input, the returned list
// can include www.example.com and the overlapping wildcard alternate domain name
// (*.example.com), if they exist. If you provide *.example.com as input, the
// returned list can include *.example.com and any alternate domain names covered
// by that wildcard (for example, www.example.com, test.example.com,
// dev.example.com, and so on), if they exist. To list conflicting aliases, you
// provide the alias to search and the ID of a distribution in your account that
// has an attached SSL/TLS certificate that includes the provided alias. For more
// information, including how to set up the distribution and certificate, see
// Moving an alternate domain name to a different distribution
// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/CNAMEs.html#alternate-domain-names-move)
// in the Amazon CloudFront Developer Guide. You can optionally specify the maximum
// number of items to receive in the response. If the total number of items in the
// list exceeds the maximum that you specify, or the default maximum, the response
// is paginated. To get the next page of items, send a subsequent request that
// specifies the NextMarker value from the current response as the Marker value in
// the subsequent request.
func (c *Client) ListConflictingAliases(ctx context.Context, params *ListConflictingAliasesInput, optFns ...func(*Options)) (*ListConflictingAliasesOutput, error) {
	if params == nil {
		params = &ListConflictingAliasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConflictingAliases", params, optFns, c.addOperationListConflictingAliasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConflictingAliasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConflictingAliasesInput struct {

	// The alias (also called a CNAME) to search for conflicting aliases.
	//
	// This member is required.
	Alias *string

	// The ID of a distribution in your account that has an attached SSL/TLS
	// certificate that includes the provided alias.
	//
	// This member is required.
	DistributionId *string

	// Use this field when paginating results to indicate where to begin in the list of
	// conflicting aliases. The response includes conflicting aliases in the list that
	// occur after the marker. To get the next page of the list, set this field’s value
	// to the value of NextMarker from the current page’s response.
	Marker *string

	// The maximum number of conflicting aliases that you want in the response.
	MaxItems *int32

	noSmithyDocumentSerde
}

type ListConflictingAliasesOutput struct {

	// A list of conflicting aliases.
	ConflictingAliasesList *types.ConflictingAliasesList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListConflictingAliasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListConflictingAliases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListConflictingAliases{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListConflictingAliasesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConflictingAliases(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListConflictingAliases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "ListConflictingAliases",
	}
}
