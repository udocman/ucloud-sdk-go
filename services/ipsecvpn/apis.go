// Code is generated by ucloud-model, DO NOT EDIT IT.

package ipsecvpn

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// IPSecVPN API Schema

// CreateRemoteVPNGatewayRequest is request schema for CreateRemoteVPNGateway action
type CreateRemoteVPNGatewayRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// 备注，默认为空
	Remark *string `required:"false"`

	// 客户VPN网关地址
	RemoteVPNGatewayAddr *string `required:"true"`

	// 客户VPN网关名称
	RemoteVPNGatewayName *string `required:"true"`

	// 业务组名称，默认为 "Default"
	Tag *string `required:"false"`
}

// CreateRemoteVPNGatewayResponse is response schema for CreateRemoteVPNGateway action
type CreateRemoteVPNGatewayResponse struct {
	response.CommonBase

	// 新建客户VPN网关的资源ID
	RemoteVPNGatewayId string
}

// NewCreateRemoteVPNGatewayRequest will create request of CreateRemoteVPNGateway action.
func (c *IPSecVPNClient) NewCreateRemoteVPNGatewayRequest() *CreateRemoteVPNGatewayRequest {
	req := &CreateRemoteVPNGatewayRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: CreateRemoteVPNGateway

创建客户VPN网关
*/
func (c *IPSecVPNClient) CreateRemoteVPNGateway(req *CreateRemoteVPNGatewayRequest) (*CreateRemoteVPNGatewayResponse, error) {
	var err error
	var res CreateRemoteVPNGatewayResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreateRemoteVPNGateway", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// CreateVPNGatewayRequest is request schema for CreateVPNGateway action
type CreateVPNGatewayRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// 业务组ID
	BusinessId *string `required:"false"`

	// 付费方式, 枚举值为: Year, 按年付费; Month, 按月付费；Dynamic, 按需付费(需开启权限)；Trial, 试用(需开启权限)；默认为按月付费
	ChargeType *string `required:"false"`

	// 代金券ID, 默认不使用
	CouponId *string `required:"false"`

	// 若要绑定EIP，在此填上EIP的资源ID
	EIPId *string `required:"false"`

	// 购买的VPN网关规格，枚举值为: Standard, 标准型; Enhanced, 增强型
	Grade *string `required:"true"`

	// 购买时长, 默认: 1
	Quantity *int `required:"false"`

	// 备注，默认为空
	Remark *string `required:"false"`

	// 业务组名称，默认为 "Default"
	Tag *string `required:"false"`

	// 新建VPN网关所属VPC的资源ID
	VPCId *string `required:"true"`

	// 新建VPN网关名称
	VPNGatewayName *string `required:"true"`
}

// CreateVPNGatewayResponse is response schema for CreateVPNGateway action
type CreateVPNGatewayResponse struct {
	response.CommonBase

	// 新建VPN网关的资源ID
	VPNGatewayId string
}

// NewCreateVPNGatewayRequest will create request of CreateVPNGateway action.
func (c *IPSecVPNClient) NewCreateVPNGatewayRequest() *CreateVPNGatewayRequest {
	req := &CreateVPNGatewayRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: CreateVPNGateway

创建VPN网关
*/
func (c *IPSecVPNClient) CreateVPNGateway(req *CreateVPNGatewayRequest) (*CreateVPNGatewayResponse, error) {
	var err error
	var res CreateVPNGatewayResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreateVPNGateway", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// CreateVPNTunnelRequest is request schema for CreateVPNTunnel action
type CreateVPNTunnelRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// IKE协商过程中使用的认证算法，"md5", "sha1", "sha2-256"。默认值为“sha1”
	IKEAuthenticationAlgorithm *string `required:"false"`

	// IKE协商过程中使用的DH组，枚举值，"1", "2", "5", "14", "15", "16"。默认为“15”
	IKEDhGroup *string `required:"false"`

	// IKE协商过程中使用的加密算法，枚举值，"aes128", "aes192", "aes256", "aes512", "3des"。默认值为“aes128”
	IKEEncryptionAlgorithm *string `required:"false"`

	// IKE协商过程中使用的模式，枚举值，主模式，“main”；野蛮模式，“aggressive”。IKEV1默认为主模式“main”，IKEV2时不使用该参数。
	IKEExchangeMode *string `required:"false"`

	// 本端标识。枚举值，自动识别，“auto”；IP地址或域名。默认为自动识别“auto”。IKEV2必填该参数
	IKELocalId *string `required:"false"`

	// 预共享密钥
	IKEPreSharedKey *string `required:"true"`

	// 客户端标识。枚举值，自动识别，“auto”；IP地址或域名。默认为“自动识别“auto”。IKEV2必填该参数
	IKERemoteId *string `required:"false"`

	// IKE中SA的生存时间，可填写范围为600-604800。默认为86400。
	IKESALifetime *string `required:"false"`

	// ike版本，枚举值： "IKE V1"，"IKE V2"，默认v1
	IKEVersion *string `required:"true"`

	// IPSec隧道中使用的认证算法，枚举值，"md5", "sha1"。默认值为“sha1”
	IPSecAuthenticationAlgorithm *string `required:"false"`

	// IPSec隧道中使用的加密算法，枚举值，"aes128", "aes192", "aes256", "aes512", "3des"。默认值为“aes128”
	IPSecEncryptionAlgorithm *string `required:"false"`

	// 指定VPN连接的本地子网的资源ID，最多可填写10个。
	IPSecLocalSubnetIds []string `required:"true"`

	// IPSec的PFS是否开启，枚举值，，不开启，"disable"；数字表示DH组, "1", "2", "5", "14", "15", "16"。默认为“disable”。
	IPSecPFSDhGroup *string `required:"false"`

	// 使用的安全协议，枚举值，“esp”，“ah”。默认为“esp”
	IPSecProtocol *string `required:"false"`

	// 指定VPN连接的客户网段，最多可填写20个。
	IPSecRemoteSubnets []string `required:"true"`

	// IPSec中SA的生存时间，可填写范围为1200 - 604800。默认为3600
	IPSecSALifetime *string `required:"false"`

	// IPSec中SA的生存时间（以字节计）。可选为8000 – 20000000。默认使用SA生存时间，
	IPSecSALifetimeBytes *string `required:"false"`

	// 备注，默认为空
	Remark *string `required:"false"`

	// 客户VPN网关的资源ID
	RemoteVPNGatewayId *string `required:"true"`

	// 业务组，默认为“Default”
	Tag *string `required:"false"`

	// vpcId
	VPCId *string `required:"true"`

	// VPN网关的资源ID
	VPNGatewayId *string `required:"true"`

	// VPN隧道名称
	VPNTunnelName *string `required:"true"`
}

// CreateVPNTunnelResponse is response schema for CreateVPNTunnel action
type CreateVPNTunnelResponse struct {
	response.CommonBase

	// VPN隧道的资源ID
	VPNTunnelId string
}

// NewCreateVPNTunnelRequest will create request of CreateVPNTunnel action.
func (c *IPSecVPNClient) NewCreateVPNTunnelRequest() *CreateVPNTunnelRequest {
	req := &CreateVPNTunnelRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: CreateVPNTunnel

创建VPN隧道
*/
func (c *IPSecVPNClient) CreateVPNTunnel(req *CreateVPNTunnelRequest) (*CreateVPNTunnelResponse, error) {
	var err error
	var res CreateVPNTunnelResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreateVPNTunnel", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DeleteRemoteVPNGatewayRequest is request schema for DeleteRemoteVPNGateway action
type DeleteRemoteVPNGatewayRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// 客户VPN网关的资源ID
	RemoteVPNGatewayId *string `required:"true"`
}

// DeleteRemoteVPNGatewayResponse is response schema for DeleteRemoteVPNGateway action
type DeleteRemoteVPNGatewayResponse struct {
	response.CommonBase
}

// NewDeleteRemoteVPNGatewayRequest will create request of DeleteRemoteVPNGateway action.
func (c *IPSecVPNClient) NewDeleteRemoteVPNGatewayRequest() *DeleteRemoteVPNGatewayRequest {
	req := &DeleteRemoteVPNGatewayRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DeleteRemoteVPNGateway

删除客户VPN网关
*/
func (c *IPSecVPNClient) DeleteRemoteVPNGateway(req *DeleteRemoteVPNGatewayRequest) (*DeleteRemoteVPNGatewayResponse, error) {
	var err error
	var res DeleteRemoteVPNGatewayResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DeleteRemoteVPNGateway", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DeleteVPNGatewayRequest is request schema for DeleteVPNGateway action
type DeleteVPNGatewayRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// 删除VPN时是否一并释放EIP。false，只解绑EIP不删除EIP；true，解绑并释放EIP。默认是false
	ReleaseEip *bool `required:"false"`

	// VPN网关的资源ID
	VPNGatewayId *string `required:"true"`
}

// DeleteVPNGatewayResponse is response schema for DeleteVPNGateway action
type DeleteVPNGatewayResponse struct {
	response.CommonBase
}

// NewDeleteVPNGatewayRequest will create request of DeleteVPNGateway action.
func (c *IPSecVPNClient) NewDeleteVPNGatewayRequest() *DeleteVPNGatewayRequest {
	req := &DeleteVPNGatewayRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DeleteVPNGateway

删除VPN网关
*/
func (c *IPSecVPNClient) DeleteVPNGateway(req *DeleteVPNGatewayRequest) (*DeleteVPNGatewayResponse, error) {
	var err error
	var res DeleteVPNGatewayResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DeleteVPNGateway", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DeleteVPNTunnelRequest is request schema for DeleteVPNTunnel action
type DeleteVPNTunnelRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// VPN隧道的资源ID
	VPNTunnelId *string `required:"true"`
}

// DeleteVPNTunnelResponse is response schema for DeleteVPNTunnel action
type DeleteVPNTunnelResponse struct {
	response.CommonBase
}

// NewDeleteVPNTunnelRequest will create request of DeleteVPNTunnel action.
func (c *IPSecVPNClient) NewDeleteVPNTunnelRequest() *DeleteVPNTunnelRequest {
	req := &DeleteVPNTunnelRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DeleteVPNTunnel

删除VPN隧道
*/
func (c *IPSecVPNClient) DeleteVPNTunnel(req *DeleteVPNTunnelRequest) (*DeleteVPNTunnelResponse, error) {
	var err error
	var res DeleteVPNTunnelResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DeleteVPNTunnel", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribeRemoteVPNGatewayRequest is request schema for DescribeRemoteVPNGateway action
type DescribeRemoteVPNGatewayRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"true"`

	// 数据分页值, 默认为20
	Limit *int `required:"false"`

	// 数据偏移量, 默认为0
	Offset *int `required:"false"`

	// 客户VPN网关的资源ID，例如RemoteVPNGatewayIds.0代表希望获取客户VPN网关1的信息，RemoteVPNGatewayIds.1代表客户VPN网关2，如果为空，则返回当前Region中所有客户VPN网关实例的信息
	RemoteVPNGatewayIds []string `required:"false"`

	// 业务组名称，若指定则返回业务组下所有客户VPN网关信息
	Tag *string `required:"false"`
}

// DescribeRemoteVPNGatewayResponse is response schema for DescribeRemoteVPNGateway action
type DescribeRemoteVPNGatewayResponse struct {
	response.CommonBase

	// 客户VPN网关列表, 每项参数详见 RemoteVPNGatewayDataSet
	DataSet []RemoteVPNGatewayDataSet

	// 符合条件的客户VPN网关总数
	TotalCount int
}

// NewDescribeRemoteVPNGatewayRequest will create request of DescribeRemoteVPNGateway action.
func (c *IPSecVPNClient) NewDescribeRemoteVPNGatewayRequest() *DescribeRemoteVPNGatewayRequest {
	req := &DescribeRemoteVPNGatewayRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribeRemoteVPNGateway

获取客户VPN网关信息
*/
func (c *IPSecVPNClient) DescribeRemoteVPNGateway(req *DescribeRemoteVPNGatewayRequest) (*DescribeRemoteVPNGatewayResponse, error) {
	var err error
	var res DescribeRemoteVPNGatewayResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeRemoteVPNGateway", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribeVPNGatewayRequest is request schema for DescribeVPNGateway action
type DescribeVPNGatewayRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// 数据分页值。默认为20
	Limit *int `required:"false"`

	// 数据偏移量。默认为0
	Offset *int `required:"false"`

	// 业务组名称，若指定则返回指定的业务组下的所有VPN网关的信息。
	Tag *string `required:"false"`

	// VPC的资源ID，返回指定的VPC下的所有VPN网关的信息。默认返回当前Region中所有VPN网关实例的信息
	VPCId *string `required:"false"`

	// VPN网关的资源ID，例如VPNGatewayIds.0代表希望获取VPN网关1的信息，VPNGatewayIds.1代表VPN网关2，如果为空，则返回当前Region中所有VPN网关的信息
	VPNGatewayIds []string `required:"false"`
}

// DescribeVPNGatewayResponse is response schema for DescribeVPNGateway action
type DescribeVPNGatewayResponse struct {
	response.CommonBase

	// 获取的VPN网关信息列表，每项参数详见 VPNGatewayDataSet
	DataSet []VPNGatewayDataSet

	// 满足条件的VPN网关总数
	TotalCount int
}

// NewDescribeVPNGatewayRequest will create request of DescribeVPNGateway action.
func (c *IPSecVPNClient) NewDescribeVPNGatewayRequest() *DescribeVPNGatewayRequest {
	req := &DescribeVPNGatewayRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribeVPNGateway

获取VPN网关信息
*/
func (c *IPSecVPNClient) DescribeVPNGateway(req *DescribeVPNGatewayRequest) (*DescribeVPNGatewayResponse, error) {
	var err error
	var res DescribeVPNGatewayResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeVPNGateway", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribeVPNTunnelRequest is request schema for DescribeVPNTunnel action
type DescribeVPNTunnelRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// 数据分页值, 默认为20
	Limit *int `required:"false"`

	// 数据偏移量, 默认为0
	Offset *int `required:"false"`

	// 业务组名称，若指定则返回指定的业务组下的所有VPN网关的信息
	Tag *string `required:"false"`

	// VPN隧道的资源ID，例如VPNTunnelIds.0代表希望获取信息的VPN隧道1，VPNTunneIds.1代表VPN隧道2，如果为空，则返回当前Region中所有的VPN隧道实例
	VPNTunnelIds []string `required:"false"`
}

// DescribeVPNTunnelResponse is response schema for DescribeVPNTunnel action
type DescribeVPNTunnelResponse struct {
	response.CommonBase

	// 获取的VPN隧道信息列表，每项参数详见 VPNTunnelDataSet
	DataSet []VPNTunnelDataSet

	// VPN隧道总数
	TotalCount int
}

// NewDescribeVPNTunnelRequest will create request of DescribeVPNTunnel action.
func (c *IPSecVPNClient) NewDescribeVPNTunnelRequest() *DescribeVPNTunnelRequest {
	req := &DescribeVPNTunnelRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribeVPNTunnel

获取VPN隧道信息
*/
func (c *IPSecVPNClient) DescribeVPNTunnel(req *DescribeVPNTunnelRequest) (*DescribeVPNTunnelResponse, error) {
	var err error
	var res DescribeVPNTunnelResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeVPNTunnel", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// GetVPNGatewayPriceRequest is request schema for GetVPNGatewayPrice action
type GetVPNGatewayPriceRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// 付费方式, 枚举值为: Year, 按年付费; Month, 按月付费; Dynamic, 按需付费(需开启权限); 默认为获取三种价格
	ChargeType *string `required:"false"`

	// VPN网关规格。枚举值，包括：标准型：Standard，增强型：Enhanced。
	Grade *string `required:"true"`

	// 购买时长, 默认: 1
	Quantity *int `required:"false"`
}

// GetVPNGatewayPriceResponse is response schema for GetVPNGatewayPrice action
type GetVPNGatewayPriceResponse struct {
	response.CommonBase

	// 获取的VPN网关价格信息列表，每项参数详见 VPNGatewayPriceSet
	PriceSet []VPNGatewayPriceSet
}

// NewGetVPNGatewayPriceRequest will create request of GetVPNGatewayPrice action.
func (c *IPSecVPNClient) NewGetVPNGatewayPriceRequest() *GetVPNGatewayPriceRequest {
	req := &GetVPNGatewayPriceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: GetVPNGatewayPrice

获取VPN价格
*/
func (c *IPSecVPNClient) GetVPNGatewayPrice(req *GetVPNGatewayPriceRequest) (*GetVPNGatewayPriceResponse, error) {
	var err error
	var res GetVPNGatewayPriceResponse

	reqCopier := *req

	err = c.Client.InvokeAction("GetVPNGatewayPrice", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// GetVPNGatewayUpgradePriceRequest is request schema for GetVPNGatewayUpgradePrice action
type GetVPNGatewayUpgradePriceRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// 更改的VPN网关规格，枚举值为: Standard, 标准型; Enhanced, 增强型。
	Grade *string `required:"true"`

	// VPN网关的资源ID
	VPNGatewayId *string `required:"true"`
}

// GetVPNGatewayUpgradePriceResponse is response schema for GetVPNGatewayUpgradePrice action
type GetVPNGatewayUpgradePriceResponse struct {
	response.CommonBase

	// 调整规格后的VPN网关价格, 单位为"元", 如需退费此处为负值
	Price float64
}

// NewGetVPNGatewayUpgradePriceRequest will create request of GetVPNGatewayUpgradePrice action.
func (c *IPSecVPNClient) NewGetVPNGatewayUpgradePriceRequest() *GetVPNGatewayUpgradePriceRequest {
	req := &GetVPNGatewayUpgradePriceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: GetVPNGatewayUpgradePrice

获取VPN网关规格改动价格
*/
func (c *IPSecVPNClient) GetVPNGatewayUpgradePrice(req *GetVPNGatewayUpgradePriceRequest) (*GetVPNGatewayUpgradePriceResponse, error) {
	var err error
	var res GetVPNGatewayUpgradePriceResponse

	reqCopier := *req

	err = c.Client.InvokeAction("GetVPNGatewayUpgradePrice", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// UpdateVPNGatewayRequest is request schema for UpdateVPNGateway action
type UpdateVPNGatewayRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// 网关规格。枚举值为: Standard, 标准型; Enhanced, 增强型。
	Grade *string `required:"true"`

	// VPN网关的资源ID
	VPNGatewayId *string `required:"true"`
}

// UpdateVPNGatewayResponse is response schema for UpdateVPNGateway action
type UpdateVPNGatewayResponse struct {
	response.CommonBase
}

// NewUpdateVPNGatewayRequest will create request of UpdateVPNGateway action.
func (c *IPSecVPNClient) NewUpdateVPNGatewayRequest() *UpdateVPNGatewayRequest {
	req := &UpdateVPNGatewayRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: UpdateVPNGateway

更改VPN网关规格
*/
func (c *IPSecVPNClient) UpdateVPNGateway(req *UpdateVPNGatewayRequest) (*UpdateVPNGatewayResponse, error) {
	var err error
	var res UpdateVPNGatewayResponse

	reqCopier := *req

	err = c.Client.InvokeAction("UpdateVPNGateway", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// UpdateVPNTunnelAttributeRequest is request schema for UpdateVPNTunnelAttribute action
type UpdateVPNTunnelAttributeRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// IKE协商过程中使用的认证算法
	IKEAuthenticationAlgorithm *string `required:"false"`

	// IKE协商过程中使用的DH组
	IKEDhGroup *string `required:"false"`

	// IKE协商过程中使用的加密算法
	IKEEncryptionAlgorithm *string `required:"false"`

	// IKE协商过程中使用的模式，可选“主动模式”与“野蛮模式”。IKEV2不使用该参数。
	IKEExchangeMode *string `required:"false"`

	// 本端标识。不填时默认使用之前的参数，结合IKEversion进行校验，IKEV2时不能为auto。
	IKELocalId *string `required:"false"`

	// 预共享密钥
	IKEPreSharedKey *string `required:"false"`

	// 客户端标识。不填时默认使用之前的参数，结合IKEversion进行校验，IKEV2时不能为auto。
	IKERemoteId *string `required:"false"`

	// IKE中SA的生存时间
	IKESALifetime *string `required:"false"`

	// 枚举值："IKE V1","IKE V2"
	IKEVersion *string `required:"false"`

	// IPSec隧道中使用的认证算法
	IPSecAuthenticationAlgorithm *string `required:"false"`

	// IPSec隧道中使用的加密算法
	IPSecEncryptionAlgorithm *string `required:"false"`

	// 指定VPN连接的本地子网的id，用逗号分隔
	IPSecLocalSubnetIds []string `required:"false"`

	// IPSec中的PFS是否开启
	IPSecPFSDhGroup *string `required:"false"`

	// 使用的安全协议，ESP或AH
	IPSecProtocol *string `required:"false"`

	// 指定VPN连接的客户网段，用逗号分隔
	IPSecRemoteSubnets []string `required:"false"`

	// IPSec中SA的生存时间
	IPSecSALifetime *string `required:"false"`

	// IPSec中SA的生存时间（以字节计）
	IPSecSALifetimeBytes *string `required:"false"`

	// VPN隧道的资源ID
	VPNTunnelId *string `required:"true"`
}

// UpdateVPNTunnelAttributeResponse is response schema for UpdateVPNTunnelAttribute action
type UpdateVPNTunnelAttributeResponse struct {
	response.CommonBase
}

// NewUpdateVPNTunnelAttributeRequest will create request of UpdateVPNTunnelAttribute action.
func (c *IPSecVPNClient) NewUpdateVPNTunnelAttributeRequest() *UpdateVPNTunnelAttributeRequest {
	req := &UpdateVPNTunnelAttributeRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: UpdateVPNTunnelAttribute

更新VPN隧道属性
*/
func (c *IPSecVPNClient) UpdateVPNTunnelAttribute(req *UpdateVPNTunnelAttributeRequest) (*UpdateVPNTunnelAttributeResponse, error) {
	var err error
	var res UpdateVPNTunnelAttributeResponse

	reqCopier := *req

	err = c.Client.InvokeAction("UpdateVPNTunnelAttribute", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}