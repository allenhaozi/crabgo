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
	"os"

	hydraclient "github.com/ory/hydra-client-go/client"
	kratos "github.com/ory/kratos-client-go"
	"github.com/pkg/errors"
)

var (
	defaultWebPort = "9000"
)

type IAMConfig struct {
	*LogConfig
	WebServerPort     string
	HydraPublicHost   string
	KratosPublicHost  string
	HydraPublicClient *hydraclient.OryHydra
	HydraAdminClient  *hydraclient.OryHydra
	KratosAPIClient   *kratos.APIClient
}

func GetIAMConfig() (*IAMConfig, error) {
	cfg := &IAMConfig{}
	err := cfg.initConfig()
	if err != nil {
		return nil, errors.Wrap(err, "initialize configuration failed")
	}
	return cfg, nil
}

func (cfg *IAMConfig) initConfig() error {

	// initialize log config
	if err := cfg.initLogConfig(); err != nil {
		return err
	}

	if p, ok := os.LookupEnv("IAM_SERVER_PORT"); ok {
		cfg.WebServerPort = p
	} else {
		cfg.WebServerPort = defaultWebPort
	}

	// initialize configure
	cfg.HydraAdminClient = GetHydraAdminClient()
	cfg.HydraPublicClient = GetHydraPublicClient()
	// TODO: upgrade to configurable
	cfg.HydraPublicHost = "http://hydra.dev.openaios.4pd.io"
	cfg.KratosPublicHost = "http://kratos.dev.openaios.4pd.io"

	// kratos API Client
	if kc, err := GetKratosAPIClient(); err != nil {
		return err
	} else {
		cfg.KratosAPIClient = kc
	}

	return nil
}
