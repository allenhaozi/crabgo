/**
 * Copyright 2022 Crabgo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
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

	"github.com/allenhaozi/crabgo/apis/common"
	crabcorev1 "github.com/allenhaozi/crabgo/apis/core/v1"
	"github.com/allenhaozi/crabgo/pkg/register"
	"github.com/allenhaozi/crabgo/pkg/restapi/internal"
	"github.com/allenhaozi/crabgo/pkg/service/handler"
)

type DemoAPI struct {
	internal.ApiController
	*handler.DemoHandler
}

func SetupDemoAPI(cfg *register.Config, carrier *sync.Map) (string, []common.RestAPIMeta) {
	api := NewDemoAPI(cfg)
	// handler name
	h := reflect.TypeOf(api).String()
	carrier.Store(h, api)
	return h, []common.RestAPIMeta{
		common.RestAPIMeta{}.Gen(crabcorev1.GroupVersion, common.ApiGroupApis, h, http.MethodGet, "GetUser", "/users/:id"),
	}
}

func NewDemoAPI(cfg *register.Config) *DemoAPI {
	ins := &DemoAPI{}
	ins.Initial(cfg)
	ins.DemoHandler = handler.NewDemoHandler(cfg)
	return ins
}

func (d *DemoAPI) GetUserAction(ctx register.Context) register.CrabResponseIf {
	userId := ctx.Param("id")
	return d.GetUser(ctx, userId)
}
