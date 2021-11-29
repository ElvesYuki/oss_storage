package oss

// BaseObject 对象基础属性
type BaseObject struct {
	UploadStatus int    `json:"uploadStatus"` // 上传状态
	Url          string `json:"url"`          // 真实对象地址
	Bucket       string `json:"bucket"`       // 对象存储bucket
	Object       string `json:"object"`       // 对象路径
	FileName     string `json:"fileName"`     // 对象源文件name
	Format       string `json:"format"`       // 对象格式，文件名后缀
	ContentType  string `json:"contentType"`  // 对象contentType
	Size         int64  `json:"size"`         // 对象大小，单位Byte
}

// ImgObject 图像对象
type ImgObject struct {
	BaseObject
	Width     int64      `json:"width"`     // 图片宽度
	Height    int64      `json:"height"`    // 图片高度
	Thumbnail *ImgObject `json:"thumbnail"` // 缩略图
}

// VideoObject 视频对象
type VideoObject struct {
	BaseObject
	Width     int64      `json:"width"`     // 视频宽度
	Height    int64      `json:"height"`    // 视频高度
	Duration  int64      `json:"duration"`  // 视频时长  单位：毫秒
	BitRate   int64      `json:"bitRate"`   // 视频比特率 比特率，单位：Kb/s  指视频每秒传送（包含）的比特数
	Encoder   string     `json:"encoder"`   // 编码器
	FrameRate int64      `json:"frameRate"` // 帧率，单位：FPS（Frame Per Second） 指视频每秒包含的帧数
	Thumbnail *ImgObject `json:"thumbnail"` // 视频封面
}

// AudioObject 音频对象
type AudioObject struct {
	BaseObject
	Duration int64  `json:"duration"` // 音频时长
	BitRate  int64  `json:"bitRate"`  // 音频比特率 比特率，单位：Kb/s  指音频每秒传送（包含）的比特数
	Encoder  string `json:"encoder"`  // 编码器
}

// JsonObject Json对象
type JsonObject struct {
	BaseObject
	ContentExcerpt string `json:"contentExcerpt"` // 文本节选
}

// HtmlObject Html对象
type HtmlObject struct {
	BaseObject
	ContentExcerpt string   `json:"contentExcerpt"` // 文本节选
	PictureCount   int64    `json:"pictureCount"`   // 图片数量
	PictureList    []string `json:"pictureList"`    // 图片url节选
}
