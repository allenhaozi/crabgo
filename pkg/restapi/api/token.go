/**
 * Copyright 2022 4Paradigm
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
	"strings"
	"sync"

	"gitlab.4pd.io/openaios/openaios-iam/apis/common"
	iamcorev1alpha1 "gitlab.4pd.io/openaios/openaios-iam/apis/core/v1alpha1"
	"gitlab.4pd.io/openaios/openaios-iam/pkg/register"
	"gitlab.4pd.io/openaios/openaios-iam/pkg/restapi/internal"
	"gitlab.4pd.io/openaios/openaios-iam/pkg/service/handler"
)

type TokenAPI struct {
	internal.ApiController
	*handler.TokenHandler
}

func SetupTokenAPI(cfg *register.Config, carrier *sync.Map) (string, []common.RestAPIMeta) {
	api := NewTokenAPI(cfg)
	// handler name
	h := reflect.TypeOf(api).String()
	carrier.Store(h, api)

	return h, []common.RestAPIMeta{
		common.RestAPIMeta{}.Gen(iamcorev1alpha1.GroupVersion, common.ApiGroupApis, h, http.MethodGet, "Oauth2Token", "/oauth2/token"),
		common.RestAPIMeta{}.Gen(iamcorev1alpha1.GroupVersion, common.ApiGroupApis, h, http.MethodGet, "UserInfo", "/userinfo"),
	}
}

func NewTokenAPI(cfg *register.Config) *TokenAPI {
	ins := &TokenAPI{}
	ins.Initial(cfg)
	ins.TokenHandler = handler.NewTokenHandler(cfg)

	return ins
}

func (t *TokenAPI) Oauth2TokenAction(ctx register.Context) register.IAMResponseIf {
	code := ctx.QueryParam("code")
	params := &iamcorev1alpha1.OAuth2AuthParams{}

	responseType := ctx.QueryParam("response_type")
	redirectUrl := ctx.QueryParam("redirect_uri")
	scope := ctx.QueryParam("scope")

	// compose oauth2 parameters
	params.ID = ctx.QueryParam("client_id")
	params.State = ctx.QueryParam("state")

	params.RedirectURIs = []string{redirectUrl}
	params.ResponseTypes = []string{responseType}
	params.Scopes = []string{scope}

	// only support basicOAuth
	basicOAuth := ctx.Request().Header.Get("Authorization")

	return t.TokenHandler.Oauth2Token(ctx, code, basicOAuth, params)

}

func (t *TokenAPI) UserInfoAction(ctx register.Context) register.IAMResponseIf {

	authInfo := ctx.Request().Header.Get("Authorization")

	var token string

	// parse Authorization Bearer Token
	tokens := strings.Split(authInfo, " ")

	if len(token) == 2 {
		token = tokens[1]
	}

	return t.TokenHandler.GetUserInfo(ctx, token)

}
