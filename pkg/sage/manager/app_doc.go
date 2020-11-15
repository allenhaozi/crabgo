package manager

import (
	"github.com/allenhaozi/crabgo/pkg/register"
	appdocrender "github.com/allenhaozi/crabgo/pkg/sage/render/app_doc"
	"github.com/allenhaozi/crabgo/pkg/sage/services"
)

type SageAppDocManager struct {
	appDocSvc *services.AppDocService
	sageBuilder  *appdocrender.SageAppDocRender
}

func NewAppDocManagerManager(cfg *register.Config) *SageAppDocManager {
	nam := &SageAppDocManager{
		appDocSvc: services.NewAppDocService(cfg),
	}
	return nam
}
