package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

type LabelTool struct {
}

func (lt LabelTool) GenLabelIndexs(labels map[string]string) string {

	if len(labels) <= 0 {
		return ""
	}

	var md5List []string
	for k, v := range labels {
		keyMd5 := lt.GetDigestMd5(k)
		md5List = append(md5List, keyMd5)

		kv := fmt.Sprintf("%s:%s", k, v)
		kvMd5 := lt.GetDigestMd5(kv)
		md5List = append(md5List, kvMd5)
	}

	labelIndexes := strings.Join(md5List, " ")
	return labelIndexes
}

func (lt LabelTool) GenLabelSearchValue(labels map[string]string) string {

	var md5List []string
	md5List = append(md5List, "")

	for k, v := range labels {
		var kv string
		if len(v) <= 0 {
			kv = k
		} else {
			kv = fmt.Sprintf("%s:%s", k, v)
		}

		keyMd5 := lt.GetDigestMd5(kv)
		md5List = append(md5List, keyMd5)
	}

	searchValue := strings.Join(md5List, " +")
	return searchValue
}

func (lt LabelTool) GetDigestMd5(key string) string {
	data := []byte(key)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func (lt LabelTool) ParseStrLabels(labels string) (map[string]string, error) {
	// if current labels not null, merge the label
	curLabelsMap := map[string]string{}
	// if labels empty, return a empty map
	if len(labels) <= 0 {
		return curLabelsMap, nil
	}

	err := json.Unmarshal([]byte(labels), &curLabelsMap)
	if err != nil {
		return curLabelsMap, errors.Wrap(err, "parse string labels failed")
	}

	return curLabelsMap, nil
}

func (lt LabelTool) CheckLabelExist(curLabelsMap, appendLabelMap map[string]string) bool {

	for k, v := range appendLabelMap {
		if vv, ok := curLabelsMap[k]; ok {
			// key same but value diff
			if vv != v {
				return false
			}
		} else {
			// key diff
			return false
		}
	}

	return true
}

func (lt LabelTool) MergeLabels(curLabelsMap map[string]string, appendLabelMap map[string]string) (string, map[string]string, error) {
	// if labels have the data merge
	for k, v := range appendLabelMap {
		// if k exist, overwrite
		curLabelsMap[k] = v
	}

	b, err := json.Marshal(curLabelsMap)

	if err != nil {
		return "", nil, errors.Wrap(err, "merge labels failed")
	}
	return string(b), curLabelsMap, err
}
