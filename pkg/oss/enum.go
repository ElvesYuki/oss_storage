package oss

var objectTypeMap map[string]*ObjectTypeItem

type ObjectType struct {
	OBJECT_TYPE_DEFAULT *ObjectTypeItem

	OBJECT_TYPE_IMAGE *ObjectTypeItem
	OBJECT_TYPE_VIDEO *ObjectTypeItem
	OBJECT_TYPE_AUDIO *ObjectTypeItem

	OBJECT_TYPE_JSON *ObjectTypeItem
	OBJECT_TYPE_HTML *ObjectTypeItem
}

var objectTypeEnum = &ObjectType{
	OBJECT_TYPE_DEFAULT: newObjectTypeItem("default", "default", "default", "默认类型，需要自行判断"),
	//
	OBJECT_TYPE_IMAGE: newObjectTypeItem("image", "default", "default", "图像类型"),
	OBJECT_TYPE_VIDEO: newObjectTypeItem("video", "default", "default", "视频类型"),
	OBJECT_TYPE_AUDIO: newObjectTypeItem("audio", "default", "default", "音频类型"),

	OBJECT_TYPE_JSON: newObjectTypeItem("json", "json", "application/json", "上传类型为json字符串，做文本存储用"),
	OBJECT_TYPE_HTML: newObjectTypeItem("html", "html", "text/html", "富文本片段"),
}

type ObjectTypeItem struct {
	ObjectType   string
	ObjectSuffix string
	ContentType  string
	Desc         string
}

func newObjectTypeItem(ObjectType string, ObjectSuffix string, ContentType string, Desc string) *ObjectTypeItem {
	return &ObjectTypeItem{
		ObjectType:   ObjectType,
		ObjectSuffix: ObjectSuffix,
		ContentType:  ContentType,
		Desc:         Desc,
	}
}
