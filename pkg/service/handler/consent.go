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

	log "github.com/sirupsen/logrus"
	"gitlab.4pd.io/openaios/openaios-iam/pkg/register"
	"gitlab.4pd.io/openaios/openaios-iam/pkg/service/mesh"
)

type ConsentHandler struct {
	cfg            *register.Config
	consentService *mesh.ConsentService
}

func NewConsentHandler(cfg *register.Config) *ConsentHandler {
	return &ConsentHandler{cfg: cfg, consentService: mesh.NewConsentService(cfg)}
}

func (c *ConsentHandler) Consent(ctx register.Context, consentChallenge string) register.IAMResponseIf {

	log.Infof("Consent request, consent challenge: %s", consentChallenge)
	errInfo := register.NewCrabError()

	if len(consentChallenge) == 0 {
		errInfo.SetData(http.StatusBadRequest, register.IAMResponseTypeJSON, "invalid consent challenge")
		return errInfo
	}

	acceptConsent, err := c.consentService.CheckConsentChallenge(ctx, consentChallenge)

	if err != nil {
		log.Errorf("Consent request failed, err: %v", err)
		errInfo.SetData(http.StatusBadRequest, register.IAMResponseTypeJSON, err.Error())
		return errInfo
	}

	resp := register.NewIAMResponse()

	resp.SetData(http.StatusFound, register.IAMResponseTypeRedirect, acceptConsent.GetPayload().RedirectTo)

	return resp

}
