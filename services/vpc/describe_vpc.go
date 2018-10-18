//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api VPC DescribeVPC

package vpc

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeVPCRequest is request schema for DescribeVPC action
type DescribeVPCRequest struct {
	request.CommonBase

	// VPCId
	VPCIds []string `required:"false"`

	// 业务组名称
	Tag *string `required:"false"`
}

// DescribeVPCResponse is response schema for DescribeVPC action
type DescribeVPCResponse struct {
	response.CommonBase

	// vpc信息
	DataSet []VPCInfo
}

// NewDescribeVPCRequest will create request of DescribeVPC action.
func (c *VPCClient) NewDescribeVPCRequest() *DescribeVPCRequest {
	req := &DescribeVPCRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeVPC - 获取VPC信息
func (c *VPCClient) DescribeVPC(req *DescribeVPCRequest) (*DescribeVPCResponse, error) {
	var err error
	var res DescribeVPCResponse

	err = c.client.InvokeAction("DescribeVPC", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}