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

package mesh

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	hydraadmin "github.com/ory/hydra-client-go/client/admin"
	hydramodels "github.com/ory/hydra-client-go/models"
	kratosclient "github.com/ory/kratos-client-go"
	"github.com/pkg/errors"

	iamcorev1alpha1 "github.com/allenhaozi/crabgo/apis/core/v1alpha1"
	"github.com/allenhaozi/crabgo/pkg/register"
)

type LoginService struct {
	cfg *register.Config
}

func NewLoginService(cfg *register.Config) *LoginService {
	return &LoginService{cfg: cfg}
}

func (s *LoginService) LoginAccept(ctx register.Context, loginChallenge, cookie string) (string, error) {
	c := context.Background()
	session, _, err := s.cfg.GeneralConfig.IAMConfig.KratosAPIClient.V0alpha2Api.ToSession(c).Cookie(cookie).Execute()

	// if there is a valid session, marshal session.identity.traits to json to be stored in subject
	traitsJSON, err := json.Marshal(session.Identity.Traits)
	if err != nil {
		return "", errors.Wrap(err, "marshal session identity failed")
	}
	subject := string(traitsJSON)

	loginAcceptParam := hydraadmin.NewAcceptLoginRequestParams()

	loginAcceptParam.SetContext(c)
	loginAcceptParam.SetLoginChallenge(loginChallenge)

	body := &hydramodels.AcceptLoginRequest{
		Remember:    true,
		RememberFor: 3600,
		Subject:     &subject,
	}

	loginAcceptParam.SetBody(body)

	// accept hydra login request
	res, err := s.cfg.GeneralConfig.IAMConfig.HydraAdminClient.Admin.AcceptLoginRequest(loginAcceptParam)

	if err != nil {
		return "", errors.Wrapf(err, "AcceptLogin failed, login_challenge:%s", loginChallenge)
	}

	return *res.GetPayload().RedirectTo, nil

}

func (s *LoginService) GetLoginInfo(ctx register.Context, params *iamcorev1alpha1.OAuth2AuthParams) (*iamcorev1alpha1.LoginResponse, error) {

	loginResponse := &iamcorev1alpha1.LoginResponse{}

	// login_challenge(hydra) and login flow id

	// if login_challenge and  flow are empty
	// return hydra /oauth2/auth
	if len(params.LoginChallenge) == 0 && len(params.LoginFlow.Id) == 0 {
		loginResponse.RedirectURL = s.GetRedirectURL(ctx, params)
		return loginResponse, nil
	}

	// if login_challenge is set but flow is empty
	if len(params.LoginFlow.Id) == 0 {
		loginResponse.RedirectURL = s.GetLoginFlowRedirectURL(ctx, params.LoginChallenge)
		return loginResponse, nil
	}

	// if flow is set
	if len(params.LoginFlow.Id) > 0 {
		flow, err := s.GetLoginFlow(ctx, params)
		if err != nil {
			return nil, err
		} else {
			loginResponse.LoginFlow = flow
			return loginResponse, nil
		}
	}

	return nil, errors.Errorf("GetLoginInfo request failed,params:%v", params)
}

func (s *LoginService) GetLoginFlow(ctx register.Context, params *iamcorev1alpha1.OAuth2AuthParams) (*kratosclient.SelfServiceLoginFlow, error) {
	c := context.Background()

	flow, _, err := s.cfg.GeneralConfig.IAMConfig.KratosAPIClient.V0alpha2Api.GetSelfServiceLoginFlow(c).Id(params.LoginFlow.Id).Cookie(params.Cookie).Execute()
	if err != nil {
		return nil, errors.Wrapf(err, "GetSelfServiceLoginFlow failed,flow:%s", params.LoginFlow.Id)
	}

	return flow, nil
}

// through this redirect url, login page can get login flow id
func (s *LoginService) GetLoginFlowRedirectURL(ctx register.Context, loginChallenge string) string {
	// http://kratos.dev.openaios.4pd.io/self-service/login/browser
	// refresh=true
	// return_to=%2Fiam-web%2Flogin%3Flogin_challenge%3Dc98ae927baf8428493d51f0350a347e2

	// build return_to url with hydra login challenge as url query parameter
	returnToParams := url.Values{
		"login_challenge": []string{loginChallenge},
	}
	returnTo := "/iam-ui/login?" + returnToParams.Encode()
	// build redirect url with return_to as url query parameter
	// refresh=true forces a new login from kratos regardless of browser sessions
	// this is important because we are letting Hydra handle sessions
	redirectToParam := url.Values{
		"return_to": []string{returnTo},
		"refresh":   []string{"true"},
	}

	redirectTo := fmt.Sprintf("%s/self-service/login/browser?", s.cfg.GeneralConfig.IAMConfig.KratosPublicHost) + redirectToParam.Encode()

	return redirectTo

}

func (s *LoginService) GetRedirectURL(ctx register.Context, params *iamcorev1alpha1.OAuth2AuthParams) string {

	query := url.Values{
		"response_type": params.GetResponseTypes(),
		//"prompt":        []string{prompt},
		//"refresh_type":  []string{params.},
		"client_id":    []string{params.GetID()},
		"scope":        params.GetScopes(),
		"redirect_uri": params.GetRedirectURIs(),
		"state":        []string{params.State},
	}
	// https://www.ory.sh/hydra/docs/reference/api#operation/oauthAuth
	// compose redirect to hydra `/oauth2/auth` url with parameters
	redirectTo := fmt.Sprintf("%s/oauth2/auth?", s.cfg.GeneralConfig.IAMConfig.HydraPublicHost) + query.Encode()

	return redirectTo
}

func (s *LoginService) Logout(ctx register.Context, cookie string) (string, error) {
	flow, _, err := s.cfg.GeneralConfig.IAMConfig.KratosAPIClient.V0alpha2Api.CreateSelfServiceLogoutFlowUrlForBrowsers(context.Background()).Cookie(cookie).Execute()

	if err != nil {
		return "", errors.Wrap(err, "create logout flow failed")
	}

	return flow.LogoutUrl, nil
}
