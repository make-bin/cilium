// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetManagedPrefixListEntriesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`

	// The ID of the prefix list.
	//
	// PrefixListId is a required field
	PrefixListId *string `type:"string" required:"true"`

	// The version of the prefix list for which to return the entries. The default
	// is the current version.
	TargetVersion *int64 `type:"long"`
}

// String returns the string representation
func (s GetManagedPrefixListEntriesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetManagedPrefixListEntriesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetManagedPrefixListEntriesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.PrefixListId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PrefixListId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetManagedPrefixListEntriesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the prefix list entries.
	Entries []PrefixListEntry `locationName:"entrySet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetManagedPrefixListEntriesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetManagedPrefixListEntries = "GetManagedPrefixListEntries"

// GetManagedPrefixListEntriesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Gets information about the entries for a specified managed prefix list.
//
//    // Example sending a request using GetManagedPrefixListEntriesRequest.
//    req := client.GetManagedPrefixListEntriesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/GetManagedPrefixListEntries
func (c *Client) GetManagedPrefixListEntriesRequest(input *GetManagedPrefixListEntriesInput) GetManagedPrefixListEntriesRequest {
	op := &aws.Operation{
		Name:       opGetManagedPrefixListEntries,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetManagedPrefixListEntriesInput{}
	}

	req := c.newRequest(op, input, &GetManagedPrefixListEntriesOutput{})

	return GetManagedPrefixListEntriesRequest{Request: req, Input: input, Copy: c.GetManagedPrefixListEntriesRequest}
}

// GetManagedPrefixListEntriesRequest is the request type for the
// GetManagedPrefixListEntries API operation.
type GetManagedPrefixListEntriesRequest struct {
	*aws.Request
	Input *GetManagedPrefixListEntriesInput
	Copy  func(*GetManagedPrefixListEntriesInput) GetManagedPrefixListEntriesRequest
}

// Send marshals and sends the GetManagedPrefixListEntries API request.
func (r GetManagedPrefixListEntriesRequest) Send(ctx context.Context) (*GetManagedPrefixListEntriesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetManagedPrefixListEntriesResponse{
		GetManagedPrefixListEntriesOutput: r.Request.Data.(*GetManagedPrefixListEntriesOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetManagedPrefixListEntriesRequestPaginator returns a paginator for GetManagedPrefixListEntries.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetManagedPrefixListEntriesRequest(input)
//   p := ec2.NewGetManagedPrefixListEntriesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetManagedPrefixListEntriesPaginator(req GetManagedPrefixListEntriesRequest) GetManagedPrefixListEntriesPaginator {
	return GetManagedPrefixListEntriesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetManagedPrefixListEntriesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// GetManagedPrefixListEntriesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetManagedPrefixListEntriesPaginator struct {
	aws.Pager
}

func (p *GetManagedPrefixListEntriesPaginator) CurrentPage() *GetManagedPrefixListEntriesOutput {
	return p.Pager.CurrentPage().(*GetManagedPrefixListEntriesOutput)
}

// GetManagedPrefixListEntriesResponse is the response type for the
// GetManagedPrefixListEntries API operation.
type GetManagedPrefixListEntriesResponse struct {
	*GetManagedPrefixListEntriesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetManagedPrefixListEntries request.
func (r *GetManagedPrefixListEntriesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
