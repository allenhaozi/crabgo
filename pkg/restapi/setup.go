/*
 * Copyright 2021 Crabgo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package restapi

import (
	"net/http"
	"reflect"
	"sync"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/allenhaozi/crabgo/apis/common"
	"github.com/allenhaozi/crabgo/pkg/register"
	crabapi "github.com/allenhaozi/crabgo/pkg/restapi/api"
)

func Setup(cfg *register.Config) map[schema.GroupVersion][]common.RestAPIMeta {
	var restfulMods sync.Map

	list := map[schema.GroupVersion][]common.RestAPIMeta{}

	//
	// new rest api setup function should add here for registry
	//
	for _, setup := range []func(cfg *register.Config, carrier *sync.Map) (string, []common.RestAPIMeta){
		crabapi.SetupDemoAPI,
	} {
		modKey, modApiList := setup(cfg, &restfulMods)

		for _, apiMeta := range modApiList {
			apiMeta := apiMeta
			if ctrl, ok := restfulMods.Load(modKey); ok {
				patchHandlerFunc(ctrl, &apiMeta)
			}
			list[apiMeta.GroupVersion] = append(list[apiMeta.GroupVersion], apiMeta)
		}
	}

	return list
}

func patchHandlerFunc(c interface{}, meta *common.RestAPIMeta) *common.RestAPIMeta {
	meta.HandlerFunc = func(e echo.Context) error {
		errInfo := register.NewCrabError()

		methodName := meta.Act + common.ControllerFuncSuffix

		log.Infof("request start handler:%s,method:%s", meta.Handler, methodName)
		// TODO: concurrent request performance
		//  attention:
		//  only the current controller is copied
		//  the properties of the controller are still reusable
		ctrl := clone(c)

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
			if err, ok := res.(*register.CrabError); ok && !err.Success() {
				return e.JSON(err.HttpCode, err)
			}
		}

		in := []reflect.Value{
			reflect.ValueOf(ctx),
		}
		// execute request handler function
		respList := v.MethodByName(methodName).Call(in)

		log.Infof("request finished handler:%v,method:%v", meta.Handler, methodName)

		if len(respList) == 1 {
			respIf := respList[0].Interface()
			if resp, ok := respIf.(*register.CrabResponse); ok {
				switch resp.ResponseType {
				case register.CrabResponseTypeRedirect:
					url, err := resp.GetRedirectURL()
					// if occur error, return json format error message
					if err != nil {
						errInfo.SetData(http.StatusBadRequest, register.CrabResponseTypeJSON, err.Error())
						return e.JSON(errInfo.HttpCode, errInfo)
					}
					// redirect
					return e.Redirect(resp.HttpCode, url)
				default:
					// default format is json
					return e.JSON(errInfo.HttpCode, resp)
				}
			} else if errInfo, ok := respIf.(*register.CrabError); ok {
				return e.JSON(errInfo.HttpCode, errInfo)
			}
		}
		// error msg uniform json format response
		errInfo.SetData(http.StatusInternalServerError, register.CrabResponseTypeJSON)

		return e.JSON(
			errInfo.HttpCode,
			errInfo,
		)
	}
	return meta
}

func clone(from interface{}) interface{} {
	if from == nil {
		return nil
	}

	copy := reflect.New(reflect.TypeOf(from).Elem())

	val := reflect.ValueOf(from).Elem()
	nVal := copy.Elem()
	for i := 0; i < val.NumField(); i++ {
		nvField := nVal.Field(i)
		nvField.Set(val.Field(i))
	}

	return copy.Interface()
}
