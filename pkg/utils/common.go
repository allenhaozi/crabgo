package utils

import (
	"regexp"
	"time"
)

var re = regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`)

func NanoTime() string {
	return time.Now().Format(time.RFC3339Nano)
}

func IsURL(url string) bool {
	return re.MatchString(url)
}
