package restfulapi

import (
	"net/http"
	"reflect"
	"sync"

	"github.com/allenhaozi/crabgo/pkg/crab/manager"

	"github.com/allenhaozi/crabgo/pkg/apis/common"
	"github.com/allenhaozi/crabgo/pkg/register"
)

type appDocApi struct {
	ApiController
	*manager.AppDocManager
}

func SetupAppDoc(cfg *register.Config, carrier *sync.Map) (string, []common.RestfulApiMeta) {
	api := NewAppDocApi(cfg)
	k := reflect.TypeOf(api).Name()
	carrier.Store(k, api)
	return k, []common.RestfulApiMeta{
		common.RestfulApiMeta{}.Gen(k, http.MethodGet, "GetAppDoc", "/app/doc/:appId"),
	}
}

func NewAppDocApi(cfg *register.Config) *appDocApi {
	as := &appDocApi{}
	as.initial(cfg)
	as.AppDocManager = manager.NewAppDocManager(cfg)
	return as
}

func (as *appDocApi) GetAppDocAction(ctx register.Context) register.CrabResponseIf {
	appId := ctx.Param("appId")
	resp := as.GetAppDoc(ctx, appId)
	return resp
}
