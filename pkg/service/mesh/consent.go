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

package mesh

import (
	"context"

	hydraadmin "github.com/ory/hydra-client-go/client/admin"
	hydramodels "github.com/ory/hydra-client-go/models"
	"github.com/pkg/errors"

	"gitlab.4pd.io/openaios/openaios-iam/pkg/register"
)

type ConsentService struct {
	cfg *register.Config
}

func NewConsentService(cfg *register.Config) *ConsentService {
	return &ConsentService{cfg: cfg}
}

func (s *ConsentService) CheckConsentChallenge(ctx register.Context, consentChallenge string) (*hydraadmin.AcceptConsentRequestOK, error) {

	params := hydraadmin.NewGetConsentRequestParams()
	params.WithContext(context.Background())
	params.SetConsentChallenge(consentChallenge)

	res, err := s.cfg.GeneralConfig.IAMConfig.HydraAdminClient.Admin.GetConsentRequest(params)

	if err != nil {
		return nil, errors.Wrapf(err, "get consent request failed,consent_challenge:%s", consentChallenge)
	}

	acceptBody := &hydramodels.AcceptConsentRequest{
		GrantAccessTokenAudience: res.GetPayload().RequestedAccessTokenAudience,
		GrantScope:               res.Payload.RequestedScope,
	}

	acceptParams := hydraadmin.NewAcceptConsentRequestParams()
	acceptParams.WithContext(context.Background())
	acceptParams.SetConsentChallenge(consentChallenge)
	acceptParams.WithBody(acceptBody)

	acceptRes, err := s.cfg.GeneralConfig.IAMConfig.HydraAdminClient.Admin.AcceptConsentRequest(acceptParams)

	if err != nil {
		return nil, errors.Wrapf(err, "accept consent request failed,consent_challenge:%s", consentChallenge)
	}

	return acceptRes, nil
}
