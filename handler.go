package crabgo

import (
	"net/http"
)

type CrabHandler struct {
}

func (self CrabHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// parse the controller name and action name
	//execute requst by the relevant controller and action
	dispatcher := new(CrabDispatcher)
	dispatcher.Dispatch(rw, req)
}
