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

	"github.com/pkg/errors"
)

var (
	defaultWebPort = "8080"
)

type CrabConfig struct {
	*LogConfig
	WebServerPort string
}

func GetCrabConfig() (*CrabConfig, error) {
	cfg := &CrabConfig{}
	err := cfg.initConfig()
	if err != nil {
		return nil, errors.Wrap(err, "initialize configuration failed")
	}
	return cfg, nil
}

func (cfg *CrabConfig) initConfig() error {

	// initialize log config
	cfg.initLogConfig()

	// web server port
	if p, ok := os.LookupEnv("SERVER_PORT"); ok {
		cfg.WebServerPort = p
	} else {
		cfg.WebServerPort = defaultWebPort
	}

	return nil
}
