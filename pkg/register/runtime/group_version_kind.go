package runtime

type GroupVersionKind struct {
	Kind string
	GroupVersion
}

func (gvk GroupVersionKind) GetKind() string {
	return gvk.Kind
}
