package services

import (
	"github.com/allenhaozi/crabgo/pkg/register"
)

type AppDocService struct {
	AbstractService
}

func NewAppDocService(cfg *register.Config) *AppDocService {
	ads := &AppDocService{}
	ads.initAppDoc(cfg)
	return ads
}

