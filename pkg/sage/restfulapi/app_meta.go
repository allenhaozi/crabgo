package restfulapi

import (
	"github.com/allenhaozi/crabgo/pkg/apis/common"
	"github.com/allenhaozi/crabgo/pkg/register"
	"net/http"
	"reflect"
	"sync"
)

type appMetaApi struct {
	ApiController
}

func SetupAppMeta(cfg *register.Config, carrier *sync.Map) (string, []common.RestfulApiMeta) {
	api := NewAppMetaApi(cfg)
	k := reflect.TypeOf(api).Name()
	carrier.Store(k, api)
	return k, []common.RestfulApiMeta{
		common.RestfulApiMeta{}.Gen(k, http.MethodGet, "GetAppMeta", "/app/meta/:name/:version"),
		common.RestfulApiMeta{}.Gen(k, http.MethodDelete, "DeleteAppMeta", "/app/meta/:name/:version"),
		common.RestfulApiMeta{}.Gen(k, http.MethodPut, "UpdateAppMeta", "/app/meta/:name/:version"),
		common.RestfulApiMeta{}.Gen(k, http.MethodPost, "RegisterAppMeta", "/app/meta"),
	}
}

func NewAppMetaApi(cfg *register.Config) *appMetaApi {
	as := &appMetaApi{}
	as.initial(cfg)
	return as
}

func (as *appMetaApi) GetAppMetaAction(ctx register.Context) register.SageResponseIf {
	errInfo := register.NewSageError()
	return errInfo
}
