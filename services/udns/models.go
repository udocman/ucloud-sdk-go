// Code is generated by ucloud-model, DO NOT EDIT IT.

package udns

/*
ValueSet - RecordInfos
*/
type ValueSet struct {

	// 主机记录
	Data string

	// 是否启用
	IsEnabled int

	// 权重
	Weight int
}

/*
RecordInfo - DescribeUDNSRecord
*/
type RecordInfo struct {

	// 主机记录
	Name string

	// 域名记录资源ID
	RecordId string

	// 记录备注信息
	Remark string

	// TTL值，单位为秒
	TTL int

	// 记录类型
	Type string

	// 数值组
	ValueSet []ValueSet

	// 记录策略，标准或随机应答
	ValueType string
}

/*
VPCInfo - ZoneInfo
*/
type VPCInfo struct {

	// VPC名称
	Name string

	// VPC地址空间
	Network []string

	// VPC ID
	VPCId string

	// VPC所属项目ID
	VPCProjectId string

	// VPC类型：Normal 公有云 Hybrid 托管云
	VPCType string
}

/*
ZoneInfo - DescribeUDNSZone
*/
type ZoneInfo struct {

	// 计费类型（Dynamic、Month、Year）
	ChargeType string

	// 创建时间
	CreateTime int

	// 域名名称
	DNSZoneName string

	// 过期时间
	ExpireTime int

	// 是否开启自动续费（Yes No）
	IsAutoRenew string

	// 是否支持迭代。枚举值,"enable",支持迭代; "disable",不支持迭代
	IsRecursionEnabled string

	// 记录相关ID
	RecordInfos []string

	// 备注
	Remark string

	// 业务组
	Tag string

	// 绑定的VPC信息
	VPCInfos []VPCInfo
}