/**
 * Copyright 2021 4Paradigm
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
	iamcorev1alpha1 "github.com/allenhaozi/crabgo/apis/core/v1alpha1"
	"github.com/allenhaozi/crabgo/pkg/register"
	"github.com/allenhaozi/crabgo/pkg/restapi/internal"
	"github.com/allenhaozi/crabgo/pkg/service/handler"
)

type LoginAPI struct {
	internal.ApiController
	*handler.LoginHandler
}

func SetupLoginAPI(cfg *register.Config, carrier *sync.Map) (string, []common.RestAPIMeta) {
	api := NewLoginAPI(cfg)
	// handler name
	h := reflect.TypeOf(api).String()
	carrier.Store(h, api)
	return h, []common.RestAPIMeta{
		common.RestAPIMeta{}.Gen(iamcorev1alpha1.GroupVersion, common.ApiGroupApis, h, http.MethodGet, "GetLoginInfo", "/login/info"),
		common.RestAPIMeta{}.Gen(iamcorev1alpha1.GroupVersion, common.ApiGroupApis, h, http.MethodGet, "LoginAccept", "/login/accept"),
		common.RestAPIMeta{}.Gen(iamcorev1alpha1.GroupVersion, common.ApiGroupApis, h, http.MethodGet, "Logout", "/logout"),
	}
}

func NewLoginAPI(cfg *register.Config) *LoginAPI {
	ins := &LoginAPI{}
	ins.Initial(cfg)
	ins.LoginHandler = handler.NewLoginHandler(cfg)
	return ins
}

func (l *LoginAPI) LogoutAction(ctx register.Context) register.IAMResponseIf {

	cookie := ctx.Request().Header.Get("cookie")

	return l.LoginHandler.Logout(ctx, cookie)
}

func (l *LoginAPI) LoginAccept(ctx register.Context) register.IAMResponseIf {

	cookie := ctx.Request().Header.Get("cookie")

	loginChallenge := ctx.QueryParam("login_challenge")

	return l.LoginHandler.LoginAccept(ctx, loginChallenge, cookie)
}

func (l *LoginAPI) GetLoginInfoAction(ctx register.Context) register.IAMResponseIf {

	params := &iamcorev1alpha1.OAuth2AuthParams{}

	flow := ctx.QueryParam("flow")

	if len(flow) > 0 {
		params.LoginFlow.SetId(flow)
	}

	// get cookie
	cookie := ctx.Request().Header.Get("cookie")

	if len(cookie) > 0 {
		params.Cookie = cookie
	}

	responseType := ctx.QueryParam("response_type")
	redirectUrl := ctx.QueryParam("redirect_uri")
	scope := ctx.QueryParam("scope")

	// compose oauth2 parameters
	params.ID = ctx.QueryParam("client_id")
	params.State = ctx.QueryParam("state")

	params.RedirectURIs = []string{redirectUrl}
	params.ResponseTypes = []string{responseType}
	params.Scopes = []string{scope}

	loginChallenge := ctx.QueryParam("login_challenge")
	if len(loginChallenge) > 0 {
		params.LoginChallenge = loginChallenge
	}

	return l.LoginHandler.GetLoginInfo(ctx, params)
}
