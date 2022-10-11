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

	openapiclient "github.com/go-openapi/runtime/client"
	hydrapublic "github.com/ory/hydra-client-go/client/public"
	"github.com/pkg/errors"

	iamcorev1alpha1 "github.com/allenhaozi/crabgo/apis/core/v1alpha1"
	"github.com/allenhaozi/crabgo/pkg/register"
)

type TokenService struct {
	cfg *register.Config
}

func NewTokenService(cfg *register.Config) *TokenService {
	return &TokenService{cfg: cfg}
}

func (s *TokenService) Oauth2Token(ctx register.Context, code, basicOAuth string, params *iamcorev1alpha1.OAuth2AuthParams) (*hydrapublic.Oauth2TokenOK, error) {

	tokenParams := &hydrapublic.Oauth2TokenParams{}

	tokenParams.SetClientID(&params.ID)
	tokenParams.SetCode(&code)
	tokenParams.WithContext(context.Background())
	if len(params.GrantTypes) > 0 {
		tokenParams.SetGrantType(params.GrantTypes[0])
	}

	if len(params.RedirectURIs) > 0 {
		tokenParams.SetRedirectURI(&params.RedirectURIs[0])
	}

	// pass through basic auth info
	info := openapiclient.APIKeyAuth("Authorization", "header", basicOAuth)

	res, err := s.cfg.GeneralConfig.IAMConfig.HydraPublicClient.Public.Oauth2Token(tokenParams, info)

	if err != nil {
		return nil, errors.Wrapf(err, "Oauth2Token failed, code: %s,ClientID:%s", code, params.ID)
	}

	return res, nil
}

func (s *TokenService) GetUserInfo(ctx register.Context, token string) (*hydrapublic.UserinfoOK, error) {

	params := hydrapublic.NewUserinfoParams()

	authInfo := openapiclient.BearerToken(token)

	res, err := s.cfg.GeneralConfig.IAMConfig.HydraPublicClient.Public.Userinfo(params, authInfo)

	if err != nil {
		return nil, errors.Wrapf(err, "get UserInfo failed, err:%s", err.Error())
	}

	return res, nil
}
