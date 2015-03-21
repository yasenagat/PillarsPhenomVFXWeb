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
	UserCode       string
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
	ProjectCode    string
	UserCode       string
	Status         int
	InsertDatetime string
	UpdateDatetime string
}

// 原始素材
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
	ProjectCode           string
	UserCode              string
	Status                int
	InsertDatetime        string
	UpdateDatetime        string
}

// 用户自定义素材组
type MaterialFolder struct {
	FolderCode     string
	FolderName     string
	FatherCode     string
	LeafFlag       string
	FolderDetail   string
	ProjectCode    string
	UserCode       string
	Status         int
	InsertDatetime string
	UpdateDatetime string
}

// 用户自定义素材组数据
type MaterialFolderData struct {
	DataCode       string
	FolderCode     string
	MaterialCode   string
	ProjectCode    string
	UserCode       string
	Status         int
	InsertDatetime string
	UpdateDatetime string
}

// EDL文件解析
type EdlShot struct {
	ShotNum      string
	ShotType     string
	StartTime    string
	EndTime      string
	FromClipName string
	SourceFile   string
}

// 镜头,每个shot是material的一段
type Shot struct {
	ShotCode       string
	ProjectCode    string
	MaterialCode   string
	LibraryCode    string
	Picture        string
	ShotNum        string
	StartTime      string
	EndTime        string
	FromClipName   string
	SourceFile     string
	ShotType       string
	ShotName       string
	ShotFps        int
	Width          int
	Height         int
	ShotDetail     string
	ShotStatus     string
	EdlCode        string
	EdlName        string
	ShotFlag       string
	UserCode       string
	Status         int
	InsertDatetime string
	UpdateDatetime string
}

// 镜头制作需求
type ShotDemand struct {
	DemandCode     string
	ShotCode       string
	ProjectCode    string
	Picture        string
	DemandDetail   string
	DemandLevel    int
	UserCode       string
	Status         int
	InsertDatetime string
	UpdateDatetime string
}

// 镜头参考素材(来源于用户上传)
type ShotMaterial struct {
	MaterialCode   string
	ShotCode       string
	ProjectCode    string
	Picture        string
	MaterialName   string
	MaterialType   string
	MaterialDetail string
	MaterialPath   string
	UserCode       string
	Status         int
	InsertDatetime string
	UpdateDatetime string
}

// 镜头Note
type ShotNote struct {
	NoteCode       string
	ShotCode       string
	ProjectCode    string
	Picture        string
	NoteDetail     string
	NoteType       string
	NoteVerson     string
	UserCode       string
	Status         int
	InsertDatetime string
	UpdateDatetime string
}

// 用户自定义镜头组
type ShotFolder struct {
	FolderCode     string
	FolderName     string
	FatherCode     string
	LeafFlag       string
	FolderDetail   string
	ProjectCode    string
	UserCode       string
	Status         int
	InsertDatetime string
	UpdateDatetime string
}

// 用户自定义镜头组数据
type ShotFolderData struct {
	DataCode       string
	FolderCode     string
	ShotCode       string
	ProjectCode    string
	UserCode       string
	Status         int
	InsertDatetime string
	UpdateDatetime string
}

// 镜头分配外包商列表
type ShotVendor struct {
	VendorCode     string
	ProjectCode    string
	VendorUser     string
	VendorName     string
	VendorDetail   string
	OpenDetail     int
	OpenDemo       int
	DownMaterial   int
	UpDemo         int
	UpProduct      int
	UserCode       string
	Status         int
	InsertDatetime string
	UpdateDatetime string
}

// 分配给外包商的镜头列表
type ShotVendorData struct {
	DataCode       string
	VendorCode     string
	VendorUser     string
	ShotCode       string
	ProjectCode    string
	UserCode       string
	Status         int
	InsertDatetime string
	UpdateDatetime string
}

// 上传小样和成品(生成版本)
type ShotVersion struct {
	VersionCode    string
	ShotCode       string
	VendorUser     string
	VersionNum     int
	Picture        string
	Demo           string
	DemoDetail     string
	Product        string
	ProductDetail  string
	ProjectCode    string
	Status         int
	InsertDatetime string
	UpdateDatetime string
}

//---------------------- 华丽分割线以下内容尚未使用 -------------------------

// 存储shot（镜头）的缩略图，一个镜头可能有多个缩略图
type Thumbnail struct {
	ThumbnailCode  string
	ShotCode       string
	ThumbnailImage string
	Status         int
	InsertDatetime string
	UpdateDatetime string
}
type Relation struct {
	RelationCode   string
	ParentCode     string
	ChildCode      string
	ListName       string
	IsShot         bool
	Content        string //miaoasdf
	Status         int
	InsertDatetime string
	UpdateDatetime string
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

type Reference struct {
	ReferenceCode string
	ShotCode      string
	ReferenceType string
	MatName       string
	MatType       string
	MatContent    string
	MatUrl        string
	Status        int
}
