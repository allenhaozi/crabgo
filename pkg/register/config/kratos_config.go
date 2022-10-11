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

package config

import (
	"net/http"
	"net/http/cookiejar"

	kratos "github.com/ory/kratos-client-go"
	"github.com/pkg/errors"
)

// TODO: update to configurable
func GetKratosAPIClient() (*kratos.APIClient, error) {

	// create a new kratos client for self hosted server
	conf := kratos.NewConfiguration()
	conf.Servers = kratos.ServerConfigurations{{URL: "http://kratos.dev.openaios.4pd.io"}}
	cj, err := cookiejar.New(nil)
	if err != nil {
		return nil, errors.Wrap(err, "new cookiejar failed")
	}
	conf.HTTPClient = &http.Client{Jar: cj}

	return kratos.NewAPIClient(conf), nil
}
