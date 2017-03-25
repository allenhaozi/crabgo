package crabgo

import (
	"log"
	"net/http"
	"time"
)

var (
	readTimeout  = time.Millisecond * 1000
	writeTimeout = time.Millisecond * 1000
)

type CrabServer struct {
	http.Server
}

func (self *CrabServer) RunServer(addr string, handler http.Handler) {

	self.Addr = addr
	self.Handler = handler
	self.ReadTimeout = readTimeout
	err := self.ListenAndServe()
	if err != nil {
		log.Fatal()
	}
}
