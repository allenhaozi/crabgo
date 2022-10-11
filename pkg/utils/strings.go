package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"path"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	alog "github.com/sirupsen/logrus"
)

func UUID() string {
	id := uuid.New()
	return id.String()
}

func IsValidUUID(u string) error {
	_, err := uuid.Parse(u)
	return err
}

func GenLabelsKey(prefix, resource string) string {
	return fmt.Sprintf("%s/%s", prefix, resource)
}

func IntToLetter(i int) string {
	var s = "abcdefghijklmnopqrstuvwxyz"
	if i >= len(s) {
		i = len(s) - 1
	}
	return string(s[i])
}

// instance name must meet DNS-1123 standard
// for keep the container name global unique attach a string
//
func GenCommonName(flag string, index int) string {
	return fmt.Sprintf("%s-%d", flag, index)
}

func JoinStr(flag string, index int) string {
	return fmt.Sprintf("%s%d", flag, index)
}

func GenWorkloadInstanceName(appName string, index int) string {
	// for Compatible with older versions
	if index == 0 {
		return appName
	}
	return fmt.Sprintf("%s-%d", appName, index)
}

func InterfaceToString(v interface{}) string {
	var value string
	if reflect.TypeOf(v).Kind() == reflect.String {
		value = v.(string)
	} else {
		raw, err := json.Marshal(v)
		if err != nil {
			alog.Error("InterfaceToString err: ", err.Error())
			return ""
		}
		value = string(raw)
	}
	return value
}

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func JoinURL(base string, paths ...string) string {
	p := path.Join(paths...)
	return fmt.Sprintf("%s/%s", strings.TrimRight(base, "/"), strings.TrimLeft(p, "/"))
}

func NormalizeString(str string) string {
	re, _ := regexp.Compile("[._/]{1}")
	res := re.ReplaceAll([]byte(str), []byte("-"))
	return string(res)
}

func GetNormalizedGroupNameVersion(group, name, version string) string {
	return fmt.Sprintf("%s--%s--%s", NormalizeString(group), NormalizeString(name), NormalizeString(version))
}

func SplitMultiString(s string) []string {
	list := make([]string, 0)
	if strings.Contains(s, ",") {
		tmp := strings.Split(s, ",")
		for _, item := range tmp {
			if len(item) > 0 {
				list = append(list, item)
			}
		}
	} else {
		list = append(list, s)
	}
	return list
}

func ShuffleStrings(list []string) {
	if len(list) == 0 || len(list) == 1 {
		return
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })
}
