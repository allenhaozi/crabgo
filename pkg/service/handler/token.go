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
package handler

import (
	"net/http"

	iamcorev1alpha1 "github.com/allenhaozi/crabgo/apis/core/v1alpha1"
	"github.com/allenhaozi/crabgo/pkg/register"
	"github.com/allenhaozi/crabgo/pkg/service/mesh"
	log "github.com/sirupsen/logrus"
)

type TokenHandler struct {
	cfg          *register.Config
	tokenService *mesh.TokenService
}

func NewTokenHandler(cfg *register.Config) *TokenHandler {
	return &TokenHandler{cfg: cfg}
}

func (h *TokenHandler) Oauth2Token(ctx register.Context, code, basicOAuth string, params *iamcorev1alpha1.OAuth2AuthParams) register.IAMResponseIf {
	// client id info

	res, err := h.tokenService.Oauth2Token(ctx, code, basicOAuth, params)

	if err != nil {
		log.Errorf("Oauth2Token failed, err:%s", err.Error())
		errInfo := register.NewCrabError()
		errInfo.SetData(http.StatusBadRequest, register.IAMResponseTypeJSON, err.Error())
		return errInfo
	}

	resp := register.NewIAMResponse()
	resp.SetData(http.StatusOK, register.IAMResponseTypeJSON, res)
	return resp
}

func (l *TokenHandler) GetUserInfo(ctx register.Context, token string) register.IAMResponseIf {

	userInfo, err := l.tokenService.GetUserInfo(ctx, token)

	if err != nil {
		log.Errorf("UserInfo failed, err:%s", err.Error())
		errInfo := register.NewCrabError()
		errInfo.SetData(http.StatusBadRequest, register.IAMResponseTypeJSON, err.Error())
		return errInfo
	}

	resp := register.NewIAMResponse()
	resp.SetData(http.StatusOK, register.IAMResponseTypeJSON, userInfo)
	return resp

}
