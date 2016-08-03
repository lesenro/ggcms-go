package models

type UpInfo struct {
	InputId, Realname, Name string
	Assignment              bool
}
type UpInfos struct {
	UpinfoList []UpInfo
}

func AttachmentToUpInfo(atta *GgcmsArticleAttr) UpInfo {
	return UpInfo{Realname: atta.AttrUrl, Name: atta.Originalname}
}
