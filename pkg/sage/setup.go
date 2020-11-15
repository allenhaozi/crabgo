/*
services controller restful api
http handler register
*/
package sage

import (
	"net/http"
	"reflect"
	"sync"

	"github.com/pkg/errors"

	"github.com/allenhaozi/crabgo/pkg/apis/common"
	sageruntime "github.com/allenhaozi/crabgo/pkg/register/runtime"
	"github.com/allenhaozi/crabgo/pkg/sage/restfulapi"

	"github.com/labstack/echo/v4"
	alog "github.com/sirupsen/logrus"

	sagecorev1 "github.com/allenhaozi/crabgo/pkg/apis/core/v1"
	"github.com/allenhaozi/crabgo/pkg/register"
)

func Setup(cfg *register.Config) map[sageruntime.GroupVersion][]common.RestfulApiMeta {

	var restfulMods sync.Map

	list := map[sageruntime.GroupVersion][]common.RestfulApiMeta{
		sagecorev1.SchemeGroupVersion: []common.RestfulApiMeta{},
	}

	for _, setup := range []func(cfg *register.Config, carrier *sync.Map) (string, []common.RestfulApiMeta){
		restfulapi.SetupAppMeta,
	} {
		modKey, modApiList := setup(cfg, &restfulMods)

		for _, apiMeta := range modApiList {
			apiMeta := apiMeta
			if ctrl, ok := restfulMods.Load(modKey); ok {
				patchHandlerFunc(ctrl, &apiMeta)
			}
			list[sagecorev1.SchemeGroupVersion] = append(list[sagecorev1.SchemeGroupVersion], apiMeta)
		}
	}

	return list
}
func patchHandlerFunc(ctrl interface{}, meta *common.RestfulApiMeta) *common.RestfulApiMeta {

	meta.HandlerFunc = func(e echo.Context) error {
		errInfo := register.NewSageError()
		// TODO make sure controller can not set struct attribute
		// it may be lead to concurrent request contamination
		v := reflect.ValueOf(ctrl)
		// initialize request context value
		// every request will trigger initialize
		ctx := register.Context{}
		ctx.Context = e
		preInValue := []reflect.Value{
			reflect.ValueOf(&ctx),
		}
		resList := v.MethodByName(common.ControllerPreExecute).Call(preInValue)
		if len(resList) == 1 {
			res := resList[0].Interface()
			if err, ok := res.(*register.SageError); ok && err.Succ() == false {
				return e.JSON(err.HttpCode, err)
			}
		}

		in := []reflect.Value{
			reflect.ValueOf(ctx),
		}
		// execute request handler function
		methodName := meta.Act + common.ControllerFuncSuffix

		if !v.MethodByName(methodName).IsValid() {
			return e.JSON(http.StatusBadRequest, errors.Errorf("invalid method name %s", methodName))
		}

		respList := v.MethodByName(methodName).Call(in)
		alog.Infof("request accept controller:%v,method:%v", meta.Mod, methodName)
		if len(respList) == 1 {
			respIf := respList[0].Interface()
			if resp, ok := respIf.(*register.SageResponse); ok {
				return e.JSON(resp.HttpCode, resp)
			} else if errInfo, ok := respIf.(*register.SageError); ok {
				return e.JSON(errInfo.HttpCode, errInfo)
			}
		}
		errInfo.SetData(http.StatusInternalServerError)

		return e.JSON(
			errInfo.HttpCode,
			errInfo,
		)
	}
	return meta
}
