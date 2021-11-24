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
	Size         string `json:"size"`         // 对象大小，单位Byte
}

type ImgObject struct {
	BaseObject
	Width  int64 `json:"width"`  // 图片宽度
	Height int64 `json:"height"` // 图片高度
}
