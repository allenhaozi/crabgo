package common

import (
	"regexp"
)

const (
	ingressAnnotationPrefix = "^nginx.ingress.kubernetes.io"
)

var (
	ingressAnnotationRegex = regexp.MustCompile(ingressAnnotationPrefix)
)

// for service
func TargetNginxIngressAnnotation(str string) bool {
	return ingressAnnotationRegex.MatchString(str)
}
