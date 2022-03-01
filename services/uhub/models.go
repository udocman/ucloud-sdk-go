// Code is generated by ucloud-model, DO NOT EDIT IT.

package uhub

/*
TagSet - Tag详细信息
*/
type TagSet struct {

	// 镜像digest值
	Digest string

	// Tag名称
	TagName string

	// 镜像更新时间
	UpdateTime string
}

/*
RepoSet - 镜像仓库
*/
type RepoSet struct {

	// 仓库创建时间
	CreateTime string

	// 镜像仓库描述
	Description string

	// 镜像仓库是否外网可以访问，可以为ture,不可以为false
	IsOutSide string

	// 镜像仓库类型,false为私有；true为公有
	IsShared string

	// 镜像仓库名称
	RepoName string

	// 仓库更新时间
	UpdateTime string
}

/*
ImageSet - 镜像信息
*/
type ImageSet struct {

	// 创建时间
	CreateTime string

	// 镜像名称
	ImageName string

	// 最新push的Tag
	LatestTag string

	// 镜像被下载次数
	PullCount int

	// 镜像仓库名称
	RepoName string

	// 修改时间
	UpdateTime string
}