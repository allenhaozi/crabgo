package v1

type SageAppDoc struct {
	ID              uint32 `json:"id"`
	AppId           string `json:"appId"`
	Type            string `json:"type"`
	DescriptionPath string `json:"descriptionPath"`
	CTime           int64  `json:"createTime"`
	MTime           int64  `json:"modifyTime"`
}
