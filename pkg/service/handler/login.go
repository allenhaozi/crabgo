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

package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	iamcorev1alpha1 "gitlab.4pd.io/openaios/openaios-iam/apis/core/v1alpha1"
	"gitlab.4pd.io/openaios/openaios-iam/pkg/register"
	"gitlab.4pd.io/openaios/openaios-iam/pkg/service/mesh"
)

type LoginHandler struct {
	cfg          *register.Config
	loginService *mesh.LoginService
}

func NewLoginHandler(cfg *register.Config) *LoginHandler {
	return &LoginHandler{cfg: cfg, loginService: mesh.NewLoginService(cfg)}
}

func (l *LoginHandler) GetLoginInfo(ctx register.Context, params *iamcorev1alpha1.OAuth2AuthParams) register.IAMResponseIf {

	log.Infof("GetLoginInfo request, params:%v", params)
	errInfo := register.NewCrabError()

	response, err := l.loginService.GetLoginInfo(ctx, params)

	if err != nil {
		log.Errorf("GetLoginInfo request failed, err: %v,params: %v", err, params)
		errInfo.SetData(http.StatusBadRequest, register.IAMResponseTypeJSON, err.Error())
		return errInfo
	}

	resp := register.NewIAMResponse()
	resp.SetData(http.StatusOK, register.IAMResponseTypeJSON, response)

	return resp
}

func (l *LoginHandler) LoginAccept(ctx register.Context, loginChallenge, cookie string) register.IAMResponseIf {
	redirectTo, err := l.loginService.LoginAccept(ctx, loginChallenge, cookie)

	if err != nil {
		errInfo := register.NewCrabError()
		errInfo.SetData(http.StatusBadRequest, register.IAMResponseTypeJSON, err.Error())
		return errInfo
	}

	resp := register.NewIAMResponse()
	resp.SetData(http.StatusFound, register.IAMResponseTypeRedirect, redirectTo)

	return resp
}

func (l *LoginHandler) Logout(ctx register.Context, cookie string) register.IAMResponseIf {
	logoutUrl, err := l.loginService.Logout(ctx, cookie)

	if err != nil {
		log.Errorf("request logout url failed,err: %v", logoutUrl)
		errInfo := register.NewCrabError()
		errInfo.SetData(http.StatusBadRequest, register.IAMResponseTypeJSON, err.Error())
		return errInfo
	}

	resp := register.NewIAMResponse()

	resp.SetData(http.StatusFound, register.IAMResponseTypeRedirect, logoutUrl)
	return resp
}
