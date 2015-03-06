package utility

// 用户管理
type User struct {
	UserCode       string
	Password       string
	DisplayName    string
	Picture        string
	Email          string
	Phone          string
	UserAuthority  string //TODO Need to change
	FilePath       string //TODO Not Use,May be delete
	Status         int
	InsertDatetime string
	UpdateDatetime string
}

// 项目管理
type Project struct {
	ProjectCode    string
	ProjectName    string
	Picture        string
	ProjectLeader  string
	ProjectType    string
	StartDatetime  string
	EndDatetime    string
	ProjectDetail  string
	Status         int
	InsertDatetime string
	UpdateDatetime string
}

// 用户添加的库
type Library struct {
	LibraryCode    string
	LibraryName    string
	LibraryPath    string
	DpxPath        string
	JpgPath        string
	MovPath        string
	UserCode       string
	ProjectCode    string
	Status         int
	InsertDatetime string
	UpdateDatetime string
}

// 所有的Material（素材）及素材的MetaData（元数据）
type Material struct {
	LibraryCode           string
	MaterialCode          string
	MaterialName          string
	MaterialType          string
	MaterialPath          string
	VideoTrackCount       int
	Width                 int
	Height                int
	VideoAudioFramerate   float32
	TimecodeFramerate     float32
	VideoFrameCount       float32
	StartAbsoluteTimecode string
	EndAbsoluteTimecode   string
	StartEdgeTimecode     string
	EndEdgeTimecode       string
	MetaData              string
	Picture               string
	UserCode              string
	ProjectCode           string
	Status                int
	InsertDatetime        string
	UpdateDatetime        string
}

// 用户添加的素材组
type MaterialFolder struct {
	FolderCode     string
	FolderName     string
	FatherCode     string
	LeafFlag       string
	FolderDetail   string
	UserCode       string
	ProjectCode    string
	Status         int
	InsertDatetime string
	UpdateDatetime string
}

// 用户添加的素材组数据
type MaterialFolderData struct {
	DataCode       string
	FolderCode     string
	MaterialCode   string
	UserCode       string
	ProjectCode    string
	Status         int
	InsertDatetime string
	UpdateDatetime string
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
