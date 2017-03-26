package crabgo

import (
	"net/http"
	"strings"
)

const paramSplitFlag = "&"
const querySplitFlag = "="
const retMethodGet = "GET"
const retMethodPost = "POST"

type CrabParameter struct {
	RetParams map[string]string
}

func (self *CrabParameter) ParseParameter(req *http.Request) {
	if req.Method == retMethodGet {
		self.RetParams = parseGetData(req.URL.RawQuery)
	} else if req.Method == retMethodPost {
		self.RetParams = parsePostData(req)
	}
}

func parseGetData(rawQuery string) map[string]string {
	data := make(map[string]string)
	if len(rawQuery) <= 0 {
		return data
	}
	sliceQuery := strings.Split(rawQuery, paramSplitFlag)
	for _, value := range sliceQuery {
		sliceValue := strings.Split(value, querySplitFlag)
		k := sliceValue[0]
		v := sliceValue[1]
		data[k] = v
	}
	return data
}

func parsePostData(req *http.Request) map[string]string {
	data := make(map[string]string)
	err := req.ParseForm()
	if err == nil {
		form := req.Form
		for k, v := range form {
			data[k] = v[0]
		}
	}
	return data
}
