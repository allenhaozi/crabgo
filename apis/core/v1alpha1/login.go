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

package v1alpha1

import (
	oryfosite "github.com/ory/fosite"
	kratosclient "github.com/ory/kratos-client-go"
)

type LoginResponse struct {
	RedirectURL string                             `json:"redirectURL"`
	LoginFlow   *kratosclient.SelfServiceLoginFlow `json:"loginFlow,omitempty"`
}

// https://datatracker.ietf.org/doc/html/rfc6749#appendix-A
// define oauth2 authentication standard parameters
type OAuth2AuthParams struct {
	oryfosite.DefaultClient
	State          string
	LoginChallenge string
	LoginFlow      kratosclient.SelfServiceLoginFlow
	Cookie         string
}
