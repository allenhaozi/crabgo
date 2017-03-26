package crabgo

import (
	"encoding/json"
	"io"
	"net/http"
)

type ControllerInterface interface {
	IndexAction()
	Init(http.ResponseWriter, *http.Request)
	SetParam(string, string)
	GetParam(string) string
	GetAllParam() map[string]string
	SetResponse(string, interface{})
	GetResponse(string) interface{}
	SendJsonData()
}

type Controller struct {
	RetParam  map[string]string
	Response  map[string]interface{}
	Request   *http.Request
	ResWriter http.ResponseWriter
}

func (self *Controller) Init(rw http.ResponseWriter, req *http.Request) {
	self.RetParam = make(map[string]string)
	self.Response = make(map[string]interface{})
	self.Request = req
	self.ResWriter = rw
}

func (self *Controller) SendJsonData() {
	self.ResWriter.Header().Set("Content-Type", "application/json")
	self.ResWriter.Header().Add("Cache-Control", "no-store")
	data := self.Response["data"]
	js, err := json.Marshal(data)
	if err == nil {
		io.WriteString(self.ResWriter, string(js))
	}
}

func (self *Controller) SetResponse(key string, data interface{}) {
	self.Response[key] = data
}

func (self *Controller) GetResponse(key string) interface{} {
	if data, ok := self.RetParam[key]; ok {
		return data
	} else {
		return nil
	}
}

func (self *Controller) SetParam(key string, value string) {
	self.RetParam[key] = value
}

func (self *Controller) GetParam(key string) string {
	if data, ok := self.RetParam[key]; ok {
		return data
	} else {
		return ""
	}
}

func (self *Controller) GetAllParam() map[string]string {
	return self.RetParam
}
