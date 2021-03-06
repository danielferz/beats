// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteTransitGatewayVpcAttachmentRequest
type DeleteTransitGatewayVpcAttachmentInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the attachment.
	//
	// TransitGatewayAttachmentId is a required field
	TransitGatewayAttachmentId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTransitGatewayVpcAttachmentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTransitGatewayVpcAttachmentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTransitGatewayVpcAttachmentInput"}

	if s.TransitGatewayAttachmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayAttachmentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteTransitGatewayVpcAttachmentResult
type DeleteTransitGatewayVpcAttachmentOutput struct {
	_ struct{} `type:"structure"`

	// Information about the deleted VPC attachment.
	TransitGatewayVpcAttachment *TransitGatewayVpcAttachment `locationName:"transitGatewayVpcAttachment" type:"structure"`
}

// String returns the string representation
func (s DeleteTransitGatewayVpcAttachmentOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteTransitGatewayVpcAttachment = "DeleteTransitGatewayVpcAttachment"

// DeleteTransitGatewayVpcAttachmentRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified VPC attachment.
//
//    // Example sending a request using DeleteTransitGatewayVpcAttachmentRequest.
//    req := client.DeleteTransitGatewayVpcAttachmentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteTransitGatewayVpcAttachment
func (c *Client) DeleteTransitGatewayVpcAttachmentRequest(input *DeleteTransitGatewayVpcAttachmentInput) DeleteTransitGatewayVpcAttachmentRequest {
	op := &aws.Operation{
		Name:       opDeleteTransitGatewayVpcAttachment,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTransitGatewayVpcAttachmentInput{}
	}

	req := c.newRequest(op, input, &DeleteTransitGatewayVpcAttachmentOutput{})
	return DeleteTransitGatewayVpcAttachmentRequest{Request: req, Input: input, Copy: c.DeleteTransitGatewayVpcAttachmentRequest}
}

// DeleteTransitGatewayVpcAttachmentRequest is the request type for the
// DeleteTransitGatewayVpcAttachment API operation.
type DeleteTransitGatewayVpcAttachmentRequest struct {
	*aws.Request
	Input *DeleteTransitGatewayVpcAttachmentInput
	Copy  func(*DeleteTransitGatewayVpcAttachmentInput) DeleteTransitGatewayVpcAttachmentRequest
}

// Send marshals and sends the DeleteTransitGatewayVpcAttachment API request.
func (r DeleteTransitGatewayVpcAttachmentRequest) Send(ctx context.Context) (*DeleteTransitGatewayVpcAttachmentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTransitGatewayVpcAttachmentResponse{
		DeleteTransitGatewayVpcAttachmentOutput: r.Request.Data.(*DeleteTransitGatewayVpcAttachmentOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTransitGatewayVpcAttachmentResponse is the response type for the
// DeleteTransitGatewayVpcAttachment API operation.
type DeleteTransitGatewayVpcAttachmentResponse struct {
	*DeleteTransitGatewayVpcAttachmentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTransitGatewayVpcAttachment request.
func (r *DeleteTransitGatewayVpcAttachmentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
