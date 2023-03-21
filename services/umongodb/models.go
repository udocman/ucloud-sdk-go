// Code is generated by ucloud-model, DO NOT EDIT IT.

package umongodb

/*
DiskInfo - 磁盘信息
*/
type DiskInfo struct {

	// 磁盘id
	DiskId string

	// 磁盘容量单位GB
	DiskSize int
}

/*
NodeInfo - 节点信息
*/
type NodeInfo struct {

	// 节点所属副本集ID
	ClusterId string

	// DB实例创建时间
	CreateTime int

	// 副本集的Mongodb的版本
	DBVersion string

	// 数据盘信息
	DataDisk DiskInfo

	// 机型信息
	MachineType string

	// 机型信息ID
	MachineTypeId string

	// 节点ID
	NodeId string

	// 节点角色，Primary/Secondary/Arbiter/Startup等等
	NodeRole string

	// 节点类型
	NodeType string

	// 副本集/分片集群状态标记 Initing：初始化中，InitFailed：安装失败，Starting：启动中，StartFailed：启动失败，Running：运行，Stopping：关闭中，Stopped：已关闭, StopFailed：关闭失败，Deleting：删除中，Deleted：已删除，DeleteFailed：删除失败，Restarting：重启中，RestartFailed：重启失败。
	State string

	// 系统盘信息
	SysDisk DiskInfo

	// 虚拟节点ID
	VirtualClusterId string

	// 可用区
	Zone string

	// 可用区ID
	ZoneId int
}

/*
ReplicaInfo - 副本集群信息
*/
type ReplicaInfo struct {

	// 集群ID
	ClusterId string

	// 副本集创建时间
	CreateTime int

	// 副本集删除时间
	DeleteTime int

	// 隔离组ID
	IsolationGroupId string

	// 机器类型
	MachineType string

	// 机器类型Id
	MachineTypeId string

	// 副本集修改时间
	ModifyTime int

	// 副本集下的节点数量
	NodeCount int

	// 副本集下的节点信息
	NodeInfos []NodeInfo

	// 副本集ID
	ReplicaId string

	// 副本类型,ConfigRepl或者DataRepl
	ReplicaType string

	// 副本集/分片集群状态标记 Initing：初始化中，InitFailed：安装失败，Starting：启动中，StartFailed：启动失败，Running：运行，Stopping：关闭中，Stopped：已关闭, StopFailed：关闭失败，Deleting：删除中，Deleted：已删除，DeleteFailed：删除失败，Restarting：重启中，RestartFailed：重启失败。
	State string
}

/*
ClusterInfo - 集群信息
*/
type ClusterInfo struct {

	// 集群ID
	ClusterId string

	// 集群类型，ReplicaSet :副本集，SharedCluster：分片集
	ClusterType string

	// Config配置集群节点配置，分片集有效
	ConfigMachineType string

	// Config配置集群节点数量，分片集有效
	ConfigNodeCount int

	// ConfigSrv信息
	ConfigReplicaInfo ReplicaInfo

	// 副本集的访问地址
	ConnectURL string

	// DB实例创建时间
	CreateTime int

	// 副本集的Mongodb的版本
	DBVersion string

	// 数据副本信息
	DataReplicaInfos []ReplicaInfo

	// DB实例删除时间
	DeleteTime int

	// 磁盘空间(GB), 默认根据配置机型
	DiskSpace int

	// 实例名称
	InstanceName string

	// 计算规格
	MachineTypeId string

	// Mongos节点数量，分片集有效
	MongosCount int

	// Mongos节点信息
	MongosInfo []NodeInfo

	// 分片数量，分片集有效
	ShardCount int

	// 每分片节点数量，分片集有效
	ShardNodeCount int

	// 副本集/分片集群状态标记 Initing：初始化中，InitFailed：安装失败，Starting：启动中，StartFailed：启动失败，Running：运行，Stopping：关闭中，Stopped：已关闭, StopFailed：关闭失败，Deleting：删除中，Deleted：已删除，DeleteFailed：删除失败，Restarting：重启中，RestartFailed：重启失败。
	State string

	// 子网ID
	SubnetId string

	// 实例业务组
	Tag string

	// VPC的ID
	VPCId string

	// 可用区
	Zone string

	//
	ZoneId int
}

/*
BackupParam - 备份策略模型
*/
type BackupParam struct {

	// 自动备份保存份数
	AutoBackupCopies int

	// 自动备份预期时间范围 (24小时制，精确到分钟，00:00~23:59)
	AutoBackupToggleTime string

	// 自动备份预期周几 (1-7 表示 周一到周末，多个数据用','分割，eg: 3,7)
	AutoBackupToggleWeek string

	// 实例ID
	ClusterId string

	// 手动备份保存份数
	ManualBackupCopies int
}

/*
ConfigOptions - 配置选项
*/
type ConfigOptions struct {

	// 允许应用类型
	AllowedApplyType string

	// 描述
	Description string

	// 是否前端展示
	EnableDisplay bool

	// 是否可修改
	EnableModify bool

	// 是否需重启
	ForceRestart bool

	// 是否为默认选项
	IsDefaultOption bool

	// mongo版本
	MongodbVersion string

	// 默认值
	OptionDefaultValue string

	// 选项值格式
	OptionFormatType string

	// 配置选项名
	OptionName string

	// 配置选项类型 string,int,bool
	OptionValueType string

	// 配置选项值范围
	OptionValues string
}

/*
ConfigTemplateItem - 配置模板项
*/
type ConfigTemplateItem struct {

	// 配置名称
	ConfigName string

	// 配置选项
	ConfigOption ConfigOptions

	// 配置值
	ConfigValue string

	// 创建时间
	CreateTime int

	// itemId
	ItemId string

	// 修改时间
	ModifyTime int

	// 节点类型: DataNode:数据节点 | ConfigSrvNode:配置节点 | MongosNode:路由节点
	NodeType string

	// 模板ID
	TemplateId string
}

/*
BackupInfo - 备份数据模型
*/
type BackupInfo struct {

	// 备份ID
	BackupId string

	// 备份名称
	BackupName string

	// 备份数据大小
	BackupSize int

	// 备份类型
	BackupType string

	// 实例ID
	ClusterId string

	// 备份结束时间
	EndTime int

	// 备份开始时间
	StartTime int

	// 备份状态
	State string
}

/*
ConfigTemplate - 配置模板
*/
type ConfigTemplate struct {

	// 集群类型
	ClusterType string

	// 创建时间
	CreateTime int

	// 删除时间
	DeleteTime int

	// 模板描述
	Description string

	// 修改时间
	ModifyTime int

	// mongo版本
	MongodbVersion string

	// 模板ID
	TemplateId string

	// 模板名称
	TemplateName string

	// 模板类型 UsersTemplate DefaultTemplate
	TemplateType string
}

/*
MongodbMachineType -
*/
type MongodbMachineType struct {

	// cpu核数
	Cpu int

	// 配置简称  2C4G
	Description string

	// 配置分组，2m , 4m
	Group string

	// 机器类型ID o.mongo2m.medium，o.mongo2m.xlarge
	MachineTypeId string

	// 内存用量(GB)
	Memory int

	// 机器类型，N/O
	UHhostMachineType string
}

/*
MongoDBVersion -
*/
type MongoDBVersion struct {

	// MongoDB版本
	DBVersion string
}
