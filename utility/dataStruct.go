package utility

import (
	"time"
)

// 前端请求返回自定义状态代码和信息
type FeedbackMessage struct {
	FeedbackCode int
	FeedbackText string
	Data         interface{}
}

// 用户管理
type User struct {
	UserCode       string
	Password       string
	DisplayName    string
	Picture        string
	Email          string
	Phone          string
	UserAuthority  string
	FilePath       string
	Status         int
	InsertDatetime string
	UpdateDatetime string
}

// 所有的Material（素材）及素材的MetaData（元数据）
type Material struct {
	MaterialCode        string
	MaterialPath        string
	MaterialName        string
	MaterialType        string
	MaterialEncodedPath string
	Status              int
	InsertDatetime      string
	UpdateDatetime      string
}

// MongoDB，存储MaterialCode对应的Material所拥有的很多条MetaData
type MaterialMetadata struct {
	MetaDataCode   string
	MetaDataIndex  int
	MaterialCode   string
	MetaDataName   string
	MetaDataValue  string
	Status         int
	InsertDatetime time.Time
	UpdateDatetime time.Time
}

// 每个shot是material的一段，该表由EDL文件解析获得
type Shot struct {
	ShotCode       string
	MaterialCode   string
	InPoint        string
	OutPoint       string
	FieldName      string
	Status         int
	InsertDatetime string
	UpdateDatetime string
}

// 存储shot（镜头）的缩略图，一个镜头可能有多个缩略图
type Thumbnail struct {
	ThumbnailCode  string
	ShotCode       string
	ThumbnailImage string
	Status         int
	InsertDatetime string
	UpdateDatetime string
}

// 需求发给若干个接包方
type Award struct {
	AwardCode       string
	ShotCode        string
	RecieveUserCode string
	Require         string
	TargetPathBase  string
	AdditionPath    string
	Status          int
	InsertDatetime  string
	UpdateDatetime  string
}

// 接包方处理完之后上传到文件夹并在Daily里面进行记录
type Daily struct {
	DailyCode      string
	ShotCode       string
	AwardCode      string
	TargetPath     string
	Status         int
	InsertDatetime string
	UpdateDatetime string
}

// 项目管理
type Project struct {
	ProjectCode    string
	ProjectName    string
	Picture        string
	ProjectDetail  string
	ProjectLeader  string
	Status         int
	InsertDatetime string
	UpdateDatetime string
}
