// Code is generated by ucloud-model, DO NOT EDIT IT.

package ubill

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// UBill API Schema

// CreateRenewRequest is request schema for CreateRenew action
type CreateRenewRequest struct {
	request.CommonBase

	// 续费周期数[1~10]，按月计费资源可传值为0，表示续费到月底
	Quantity *int `required:"true"`

	// 需要续费资源ID
	ResourceId *string `required:"true"`
}

// CreateRenewResponse is response schema for CreateRenew action
type CreateRenewResponse struct {
	response.CommonBase

	// 订单号
	OrderNo string
}

// NewCreateRenewRequest will create request of CreateRenew action.
func (c *UBillClient) NewCreateRenewRequest() *CreateRenewRequest {
	req := &CreateRenewRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: CreateRenew

创建单个续费订单
*/
func (c *UBillClient) CreateRenew(req *CreateRenewRequest) (*CreateRenewResponse, error) {
	var err error
	var res CreateRenewResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreateRenew", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribeOrderDetailInfoRequest is request schema for DescribeOrderDetailInfo action
type DescribeOrderDetailInfoRequest struct {
	request.CommonBase

	//
	AzGroups []string `required:"false"`

	//
	BeginTime *int `required:"true"`

	//
	ChargeTypes []string `required:"false"`

	//
	EndTime *int `required:"true"`

	//
	Invoiceds []string `required:"false"`

	//
	Limit *int `required:"false"`

	//
	Offset *int `required:"false"`

	//
	OrderStates []string `required:"false"`

	//
	OrderTypes []string `required:"false"`

	//
	QueryAll *string `required:"false"`

	//
	Regions []string `required:"false"`

	//
	ResourceIds []string `required:"false"`

	//
	ResourceTypes []string `required:"false"`

	//
	Tags []string `required:"false"`

	//
	TradeNos []string `required:"false"`
}

// DescribeOrderDetailInfoResponse is response schema for DescribeOrderDetailInfo action
type DescribeOrderDetailInfoResponse struct {
	response.CommonBase

	//
	OrderInfos []OrderInfo
}

// NewDescribeOrderDetailInfoRequest will create request of DescribeOrderDetailInfo action.
func (c *UBillClient) NewDescribeOrderDetailInfoRequest() *DescribeOrderDetailInfoRequest {
	req := &DescribeOrderDetailInfoRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: DescribeOrderDetailInfo
*/
func (c *UBillClient) DescribeOrderDetailInfo(req *DescribeOrderDetailInfoRequest) (*DescribeOrderDetailInfoResponse, error) {
	var err error
	var res DescribeOrderDetailInfoResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeOrderDetailInfo", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// GetBalanceRequest is request schema for GetBalance action
type GetBalanceRequest struct {
	request.CommonBase
}

// GetBalanceResponse is response schema for GetBalance action
type GetBalanceResponse struct {
	response.CommonBase

	// 账户余额信息
	AccountInfo AccountInfo
}

// NewGetBalanceRequest will create request of GetBalance action.
func (c *UBillClient) NewGetBalanceRequest() *GetBalanceRequest {
	req := &GetBalanceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: GetBalance

获取账户余额
*/
func (c *UBillClient) GetBalance(req *GetBalanceRequest) (*GetBalanceResponse, error) {
	var err error
	var res GetBalanceResponse

	reqCopier := *req

	err = c.Client.InvokeAction("GetBalance", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// GetBillDataFileUrlRequest is request schema for GetBillDataFileUrl action
type GetBillDataFileUrlRequest struct {
	request.CommonBase

	// 此字段不推荐使用，建议使用BillingCycle.   若BillingCycle 和 BillPeriod同时存在，BillingCycle 优先
	BillPeriod *int `required:"false"`

	// 账单类型，传 0 时获取账单总览报表，传 1 获取账单明细报表
	BillType *int `required:"true"`

	// 账期(字符串格式，YYYY-MM，例如2021-08).   若BillingCycle 和 BillPeriod同时存在，BillingCycle 优先
	BillingCycle *string `required:"true"`

	// 获取账单总览报表时，账单的支付状态，传 0 时获取待支付账单，传 1 时获取已支付账单。获取账单明细报表时该参数无效
	PaidType *int `required:"false"`

	// 如需求其他语言版本的账单则使用此参数。默认中文。如 RequireVersion = "EN"，则提供英文版本账单。
	RequireVersion *string `required:"false"`

	// 文件版本，若为"v1"表示获取带有子用户信息的账单，可以为空
	Version *string `required:"false"`
}

// GetBillDataFileUrlResponse is response schema for GetBillDataFileUrl action
type GetBillDataFileUrlResponse struct {
	response.CommonBase

	// 交易账单数据下载URL
	FileUrl string

	// 生成的 URL是否有效，即有对应数据文件
	IsValid string
}

// NewGetBillDataFileUrlRequest will create request of GetBillDataFileUrl action.
func (c *UBillClient) NewGetBillDataFileUrlRequest() *GetBillDataFileUrlRequest {
	req := &GetBillDataFileUrlRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: GetBillDataFileUrl

生成账单数据文件下载的 url
*/
func (c *UBillClient) GetBillDataFileUrl(req *GetBillDataFileUrlRequest) (*GetBillDataFileUrlResponse, error) {
	var err error
	var res GetBillDataFileUrlResponse

	reqCopier := *req

	err = c.Client.InvokeAction("GetBillDataFileUrl", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// ListUBillDetailRequest is request schema for ListUBillDetail action
type ListUBillDetailRequest struct {
	request.CommonBase

	// 账期，YYYY-MM，比如2021-08，只支持2018-05之后的查询
	BillingCycle *string `required:"true"`

	// 计费方式 (筛选项, 默认全部)。枚举值：\\ > Dynamic:按时 \\ > Month:按月 \\ > Year:按年 \\ > Once:一次性按量 \\ > Used:按量 \\ > Post:后付费
	ChargeType *string `required:"false"`

	// 每页数量，默认值25，最大值：100。
	Limit *int `required:"false"`

	// 数据偏移量 (默认0)
	Offset *int `required:"false"`

	// 订单类型 (筛选项, 默认全部) 。枚举值：\\ > OT_BUY:新购 \\ > OT_RENEW:续费 \\ > OT_UPGRADE:升级 \\ > OT_REFUND:退费 \\ > OT_DOWNGRADE:降级 \\ > OT_SUSPEND:结算 \\ > OT_PAYMENT:删除资源回款 \\ > OT_POSTPAID_PAYMENT:后付费回款 \\ > OT_RECOVER:删除恢复 \\ > OT_POSTPAID_RENEW:过期续费回款
	OrderType *string `required:"false"`

	// 支付状态 (筛选项, 1:仅显示未支付订单; 2:仅显示已支付订单; 0:两者都显示)
	PaidState *int `required:"false"`

	// 项目名称 (筛选项, 默认全部)
	ProjectName *string `required:"false"`

	// 资源ID(筛选项, 默认全部)	支持多筛选，多筛选请在请求参数中添加多个字段例ResourceIds.0: uhost-bzgf1gh5，ResourceIds.1: uhost-gu1xpspa，
	ResourceIds []string `required:"false"`

	// 产品类型 (筛选项, 默认全部),支持多筛选，多筛选请在请求参数中添加多个字段。枚举值：\\ > uhost:云主机 \\ > udisk:普通云硬盘 \\ > udb:云数据库 \\ > eip:弹性IP \\ > ufile:对象存储 \\ > fortress_host:堡垒机 \\ > ufs:文件存储 \\ > waf:WEB应用防火墙 \\ > ues:弹性搜索 \\ > udisk_ssd:SSD云硬盘 \\ > rssd:RSSD云硬盘
	ResourceTypes []string `required:"false"`

	// 是否显示0元订单 (0 不显示, 1 显示, 默认0)
	ShowZero *int `required:"false"`

	// 用户邮箱，可以根据用户邮箱来进行筛选
	UserEmail *string `required:"false"`
}

// ListUBillDetailResponse is response schema for ListUBillDetail action
type ListUBillDetailResponse struct {
	response.CommonBase

	// 账单明细数组
	Items []BillDetailItem

	// 账单明细总长度
	TotalCount int
}

// NewListUBillDetailRequest will create request of ListUBillDetail action.
func (c *UBillClient) NewListUBillDetailRequest() *ListUBillDetailRequest {
	req := &ListUBillDetailRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: ListUBillDetail

获取某个账期内的所有消费。
*/
func (c *UBillClient) ListUBillDetail(req *ListUBillDetailRequest) (*ListUBillDetailResponse, error) {
	var err error
	var res ListUBillDetailResponse

	reqCopier := *req

	err = c.Client.InvokeAction("ListUBillDetail", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// ListUBillOverviewRequest is request schema for ListUBillOverview action
type ListUBillOverviewRequest struct {
	request.CommonBase

	// 账期，YYYY-MM格式，例如2022-02，只支持2018-05之后的查询
	BillingCycle *string `required:"true"`

	// 账单维度, product 按产品聚合,project 按项目聚合，user 按子账号聚合
	Dimension *string `required:"true"`

	// 是否显示已入账账单, 1 已入账, 0 待入账 (默认0 )
	HideUnpaid *int `required:"false"`
}

// ListUBillOverviewResponse is response schema for ListUBillOverview action
type ListUBillOverviewResponse struct {
	response.CommonBase

	// 账单聚合数据
	Items []BillOverviewItem

	// 账单总览数据总数
	TotalCount int

	// 已入账订单总额（已入账时显示）
	TotalPaidAmount string

	// 现金账户扣款总额	（已入账时显示）
	TotalPaidAmountReal string

	// 待入账订单总额（待入账时显示）
	TotalUnpaidAmount string
}

// NewListUBillOverviewRequest will create request of ListUBillOverview action.
func (c *UBillClient) NewListUBillOverviewRequest() *ListUBillOverviewRequest {
	req := &ListUBillOverviewRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: ListUBillOverview

账单总览。可按产品/项目/用户纬度获取某个账期内账单总览信息。
*/
func (c *UBillClient) ListUBillOverview(req *ListUBillOverviewRequest) (*ListUBillOverviewResponse, error) {
	var err error
	var res ListUBillOverviewResponse

	reqCopier := *req

	err = c.Client.InvokeAction("ListUBillOverview", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// ModifyAutoRenewFlagRequest is request schema for ModifyAutoRenewFlag action
type ModifyAutoRenewFlagRequest struct {
	request.CommonBase

	// 开关标识(TURN_ON: 打开; TURN_OFF: 关闭)
	Flag *string `required:"true"`

	// 资源ID
	ResourceId *string `required:"true"`
}

// ModifyAutoRenewFlagResponse is response schema for ModifyAutoRenewFlag action
type ModifyAutoRenewFlagResponse struct {
	response.CommonBase

	// 操作失败资源数量
	Fail int

	// 开关资源自动续费结果数组
	ResultSet []ResultSet

	// 操作成功资源数量
	Success int
}

// NewModifyAutoRenewFlagRequest will create request of ModifyAutoRenewFlag action.
func (c *UBillClient) NewModifyAutoRenewFlagRequest() *ModifyAutoRenewFlagRequest {
	req := &ModifyAutoRenewFlagRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: ModifyAutoRenewFlag

修改资源自动续费标识
*/
func (c *UBillClient) ModifyAutoRenewFlag(req *ModifyAutoRenewFlagRequest) (*ModifyAutoRenewFlagResponse, error) {
	var err error
	var res ModifyAutoRenewFlagResponse

	reqCopier := *req

	err = c.Client.InvokeAction("ModifyAutoRenewFlag", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
