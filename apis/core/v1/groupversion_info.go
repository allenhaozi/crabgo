package v1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	Group   = "crabgo"
	Version = "v1"
)

var (
	GroupVersion = schema.GroupVersion{Group: Group, Version: Version}
)
