// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The request to get an origin access identity's information.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2018-11-05/GetCloudFrontOriginAccessIdentityRequest
type GetCloudFrontOriginAccessIdentityInput struct {
	_ struct{} `type:"structure"`

	// The identity's ID.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetCloudFrontOriginAccessIdentityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCloudFrontOriginAccessIdentityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCloudFrontOriginAccessIdentityInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetCloudFrontOriginAccessIdentityInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	return nil
}

// The returned result of the corresponding request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2018-11-05/GetCloudFrontOriginAccessIdentityResult
type GetCloudFrontOriginAccessIdentityOutput struct {
	_ struct{} `type:"structure" payload:"CloudFrontOriginAccessIdentity"`

	// The origin access identity's information.
	CloudFrontOriginAccessIdentity *CloudFrontOriginAccessIdentity `type:"structure"`

	// The current version of the origin access identity's information. For example:
	// E2QWRUHAPOMQZL.
	ETag *string `location:"header" locationName:"ETag" type:"string"`
}

// String returns the string representation
func (s GetCloudFrontOriginAccessIdentityOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetCloudFrontOriginAccessIdentityOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ETag", protocol.StringValue(v), metadata)
	}
	if s.CloudFrontOriginAccessIdentity != nil {
		v := s.CloudFrontOriginAccessIdentity

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "CloudFrontOriginAccessIdentity", v, metadata)
	}
	return nil
}

const opGetCloudFrontOriginAccessIdentity = "GetCloudFrontOriginAccessIdentity2018_11_05"

// GetCloudFrontOriginAccessIdentityRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Get the information about an origin access identity.
//
//    // Example sending a request using GetCloudFrontOriginAccessIdentityRequest.
//    req := client.GetCloudFrontOriginAccessIdentityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2018-11-05/GetCloudFrontOriginAccessIdentity
func (c *Client) GetCloudFrontOriginAccessIdentityRequest(input *GetCloudFrontOriginAccessIdentityInput) GetCloudFrontOriginAccessIdentityRequest {
	op := &aws.Operation{
		Name:       opGetCloudFrontOriginAccessIdentity,
		HTTPMethod: "GET",
		HTTPPath:   "/2018-11-05/origin-access-identity/cloudfront/{Id}",
	}

	if input == nil {
		input = &GetCloudFrontOriginAccessIdentityInput{}
	}

	req := c.newRequest(op, input, &GetCloudFrontOriginAccessIdentityOutput{})
	return GetCloudFrontOriginAccessIdentityRequest{Request: req, Input: input, Copy: c.GetCloudFrontOriginAccessIdentityRequest}
}

// GetCloudFrontOriginAccessIdentityRequest is the request type for the
// GetCloudFrontOriginAccessIdentity API operation.
type GetCloudFrontOriginAccessIdentityRequest struct {
	*aws.Request
	Input *GetCloudFrontOriginAccessIdentityInput
	Copy  func(*GetCloudFrontOriginAccessIdentityInput) GetCloudFrontOriginAccessIdentityRequest
}

// Send marshals and sends the GetCloudFrontOriginAccessIdentity API request.
func (r GetCloudFrontOriginAccessIdentityRequest) Send(ctx context.Context) (*GetCloudFrontOriginAccessIdentityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCloudFrontOriginAccessIdentityResponse{
		GetCloudFrontOriginAccessIdentityOutput: r.Request.Data.(*GetCloudFrontOriginAccessIdentityOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCloudFrontOriginAccessIdentityResponse is the response type for the
// GetCloudFrontOriginAccessIdentity API operation.
type GetCloudFrontOriginAccessIdentityResponse struct {
	*GetCloudFrontOriginAccessIdentityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCloudFrontOriginAccessIdentity request.
func (r *GetCloudFrontOriginAccessIdentityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
