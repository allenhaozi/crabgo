package services

import (
	"github.com/allenhaozi/crabgo/pkg/register"
)

type AbstractService struct {
}

func (as *AbstractService) initialize(cfg *register.Config) {
	as.initAppMeta(cfg)
}

func (as *AbstractService) initAppMeta(cfg *register.Config) {
}

func (as *AbstractService) initAppDoc(cfg *register.Config) {
}


