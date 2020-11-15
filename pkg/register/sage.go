package register

type SageUser struct {
	Id       uint32 `json:"id"`
	Admin    bool   `json:"admin"`
	UserName string `username`
}

type SageHeader struct {
	WorkspaceId uint32 `json:"x-prophet-workspace-id"`
	UserToken   string `json:"user-token"`
	AccessKey   string `json:"access-key"`
}

type YarnStatus struct {
	Capacity    float64 `json:"capacity"`
	MaxCapacity float64 `json:"maxCapacity"`
	Name        string  `json:"name"`
	Status      string  `json:"status"`
	Allocated   float64 `json:"used"`
	Err         string  `json:"error"`
}

type SageYarnResource struct {
	Id         uint32     `json:"id"`
	Namespace  string     `json:"namespace"`
	Queue      string     `json:"queue"`
	User       string     `json:"user"`
	YarnStatus YarnStatus `json:"status"`
	HadoopId   string     `json:"hadoopId"`
	HadoopType string     `json:"hadoopType"`
	Alias      string     `json:"alias"`
	KeyTab     string     `json:"keytab"`
}

// keystone workspace meta
type SageWorkspace struct {
	WorkspaceId   uint32           `json:"id"`
	WorkspaceName string           `json:"workspaceName"`
	YarnResource  SageYarnResource `json:"yarnResource"`
	Namespace     string           `json:"namespace"`
}
