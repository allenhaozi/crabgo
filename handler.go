package crabgo

import (
	"io"
	"net/http"
	"reflect"
)

type CrabHandler struct {
}

func (self CrabHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	dispatcher := new(CrabDispatcher)
	dispatcher.Dispatch(req)

	parameter := new(CrabParameter)
	parameter.ParseParameter(req)

	controller := RouterList[dispatcher.ContrllerNameKey]
	controller.Init()
	for k, v := range parameter.RetParams {
		controller.SetParam(k, v)
	}

	value := reflect.ValueOf(controller)
	value.MethodByName(dispatcher.ActionName).Call(nil)

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Add("Cache-Control", "no-store")
	io.WriteString(rw, "This HTTP response has both headers before this text and trailers at the end.\n")
}
