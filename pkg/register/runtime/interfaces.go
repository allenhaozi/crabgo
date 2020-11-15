package runtime

type Object interface {
	GetGroupVersion() GroupVersion
}

type Index interface {
	ToString() string
}
