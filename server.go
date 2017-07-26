package crabgo

import (
	"log"
	"net/http"
)

type CrabServer struct {
	http.Server
}

func NewCrabServer() *CrabServer {
	return &CrabServer{}
}

func (self *CrabServer) RunServer() {
	err := self.ListenAndServe()
	if err != nil {
		log.Fatal()
	}
}
