package common

import (
	"errors"
	"fmt"
	"regexp"
)

const (

	// https://tools.ietf.org/html/rfc1123
	SageAppNameRegex = "^[a-z0-9-]{1,35}$"

	// in view of that use app instance name autogenerate service name
	// example: $service_name = ${app-instance-name}-${suffix-random-string}
	// suffix-random-string max length is: 50 - 30 = 20
	// https://tools.ietf.org/html/rfc1123
	SageNameRegex = "^[a-z]([-a-z0-9]*[a-z0-9])?$"

	SageNameMaxLength = 50

	// reference https://semver.org/
	SageVersionRegex = `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`
)

var (
	nameRegex    = regexp.MustCompile(SageNameRegex)
	nameAppRegex = regexp.MustCompile(SageAppNameRegex)
	versionRegex = regexp.MustCompile(SageVersionRegex)

	errInvalidNameFormat    = "invalid name format"
	errInvalidVersionFormat = "invalid version format, follow semver.org"
)

// for service
func ValidName(str string) error {
	if len(str) <= SageNameMaxLength && nameRegex.MatchString(str) {
		return nil
	}
	return errors.New(errInvalidNameFormat)
}

// for app meta and instance
func ValidAppName(str string) error {
	if nameAppRegex.MatchString(str) {
		return nil
	}
	return errors.New(errInvalidNameFormat)
}

// for all version format check
func ValidVersion(v string) error {
	if versionRegex.MatchString(v) {
		return nil
	}
	return errors.New(errInvalidVersionFormat)
}

var (
	// workload instance name + trait loop index
	SageSDNameFormat = "%s-%d"
	// app instance name to service name
	// instance name + random string + instance index
	SageSSDNameFormat = "%s-%s-%d"
)

// ingress name equal ingress trait name
// ingress trait name equal workload name + ingress item loop index
//      workload name equal app config name
//      one app config only have one component
func GetSDFormatName(name string, index int) string {
	return fmt.Sprintf(SageSDNameFormat, name, index)
}
