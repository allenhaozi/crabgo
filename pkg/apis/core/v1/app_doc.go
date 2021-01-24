package v1

type AppDoc struct {
	ID         uint32 `json:"id"`
	Name       string `json:"name"`
	AppId      string `json:"appId"`
	AppName    string `json:"appName"`
	AppVersion string `json:"appVersion"`
	User       string `json:"user"`
	UserId     uint32 `json:"userId"`
	Namespace  string `json:"namespace"`
	CTime      int64  `json:"createTime"`
	MTime      int64  `json:"modifyTime"`
}
