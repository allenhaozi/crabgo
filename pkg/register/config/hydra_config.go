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
	"github.com/go-openapi/strfmt"
	hydraclient "github.com/ory/hydra-client-go/client"
)

// TODO: update to configurable
func GetHydraAdminClient() *hydraclient.OryHydra {

	transport := &hydraclient.TransportConfig{
		BasePath: "/",
		Host:     "admin.hydra.dev.openaios.4pd.io",
		Schemes:  []string{"http"},
	}
	return hydraclient.NewHTTPClientWithConfig(strfmt.Default, transport)
}

// TODO: update to configurable
func GetHydraPublicClient() *hydraclient.OryHydra {
	transport := &hydraclient.TransportConfig{
		BasePath: "/",
		Host:     "hydra.dev.openaios.4pd.io",
		Schemes:  []string{"http"},
	}
	return hydraclient.NewHTTPClientWithConfig(strfmt.Default, transport)
}
