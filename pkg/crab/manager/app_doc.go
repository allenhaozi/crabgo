package manager

import (
	"net/http"

	"github.com/allenhaozi/crabgo/pkg/crab/mesh"
	"github.com/allenhaozi/crabgo/pkg/register"
)

type AppDocManager struct {
	appDocSvc *mesh.AppDocService
}

func NewAppDocManager(cfg *register.Config) *AppDocManager {
	nam := &AppDocManager{
		appDocSvc: mesh.NewAppDocService(cfg),
	}
	return nam
}

func (adm *AppDocManager) GetAppDoc(ctx register.Context, appId string) register.CrabResponseIf {

	errInfo := register.NewCrabError()
	doc, err := adm.appDocSvc.GetAppDoc(ctx, appId)

	if err != nil {
		errInfo.SetData(http.StatusNotAcceptable, err.Error())
		return errInfo
	}

	resp := register.NewCrabResponse()
	resp.SetData(http.StatusOK, doc)

	return resp
}
