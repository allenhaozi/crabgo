package utils

import (
	"regexp"

	"github.com/allenhaozi/crabgo/pkg/apis/common"
)

func ValidVersion(v string) bool {
	rr := regexp.MustCompile(common.SageVersionRegex)
	return rr.MatchString(v)
}
