package runtime

type GroupVersion struct {
	Group   string
	Version string
}

func (gv GroupVersion) GetGroupVersion() GroupVersion {
	return gv
}

func (gv GroupVersion) GetGroupVersionPath() string {
	return "/" + gv.Group + "/" + gv.Version
}
