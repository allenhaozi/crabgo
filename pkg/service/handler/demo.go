/**
 * Copyright 2021 Crabgo Authors
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

	"github.com/allenhaozi/crabgo/pkg/register"
	"github.com/allenhaozi/crabgo/pkg/service/mesh"
)

type DemoHandler struct {
	cfg         *register.Config
	demoService *mesh.DemoService
}

func NewDemoHandler(cfg *register.Config) *DemoHandler {
	return &DemoHandler{cfg: cfg, demoService: mesh.NewDemoService(cfg)}
}

func (l *DemoHandler) GetUser(ctx register.Context, userId string) register.CrabResponseIf {
	user, err := l.demoService.GetUser(ctx, userId)

	log.Infof("get request user id: %s", userId)

	if err != nil {
		errInfo := register.NewCrabError()
		errInfo.SetData(http.StatusBadRequest, register.CrabResponseTypeJSON, err.Error())
		return errInfo
	}

	response := register.NewCrabResponse()

	response.SetData(http.StatusOK, register.CrabResponseTypeJSON, user)

	return response
}
