package crabgo

import (
//"fmt"
//"net/http"
)

const (
	VERSION = "0.0.1"
)

type CrabInstance struct {
}

func (crab *CrabInstance) Run() {

	svr := new(CrabServer)
	addr := "10.125.67.57:8212"
	handler := CrabHandler{}
	svr.RunServer(addr, handler)
}
