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
	"sync"

	"gitlab.4pd.io/openaios/openaios-iam/apis/common"
	iamcorev1alpha1 "gitlab.4pd.io/openaios/openaios-iam/apis/core/v1alpha1"
	"gitlab.4pd.io/openaios/openaios-iam/pkg/register"
	"gitlab.4pd.io/openaios/openaios-iam/pkg/restapi/internal"
	"gitlab.4pd.io/openaios/openaios-iam/pkg/service/handler"
)

type ConsentAPI struct {
	internal.ApiController
	*handler.ConsentHandler
}

func SetupConsentAPI(cfg *register.Config, carrier *sync.Map) (string, []common.RestAPIMeta) {
	api := NewConsentAPI(cfg)
	// handler name
	h := reflect.TypeOf(api).String()
	carrier.Store(h, api)
	return h, []common.RestAPIMeta{
		common.RestAPIMeta{}.Gen(iamcorev1alpha1.GroupVersion, common.ApiGroupApis, h, http.MethodGet, "Consent", "/consent"),
	}
}

func NewConsentAPI(cfg *register.Config) *ConsentAPI {
	ins := &ConsentAPI{}
	ins.Initial(cfg)
	ins.ConsentHandler = handler.NewConsentHandler(cfg)
	return ins
}

func (c *ConsentAPI) ConsentAction(ctx register.Context) register.IAMResponseIf {
	consentChallenge := ctx.QueryParam("consent_challenge")
	return c.Consent(ctx, consentChallenge)
}
