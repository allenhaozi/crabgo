/**
 * Copyright 2021 Crabgo Authors
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

package register

import (
	"github.com/allenhaozi/crabgo/pkg/register/config"
)

var PConfig *Config

type ExtraConfig struct{}

type GeneralConfig struct {
	CrabConfig *config.CrabConfig
	MyConfig   *config.MysqlConfig
}

type Config struct {
	ExtraConfig   *ExtraConfig
	GeneralConfig *GeneralConfig
}

func NewConfig() (*Config, error) {
	c := &Config{}

	err := c.initConfig()
	if err != nil {
		return nil, err
	}

	// init clients which openaios-iam need
	err = c.initClient()
	if err != nil {
		return nil, err
	}

	// set a global var save Config value
	PConfig = c

	return c, nil
}

func (c *Config) initConfig() error {
	c.GeneralConfig = &GeneralConfig{}
	// mysql config
	/*
			myCfg, err := config.GetMysqlConfig()
			if err != nil {
				return err
			}
		c.GeneralConfig.MyConfig = myCfg
	*/

	cfg, err := config.GetCrabConfig()
	if err != nil {
		return err
	}

	// openaios config
	c.GeneralConfig.CrabConfig = cfg

	return nil
}

func (c *Config) initClient() error {
	return nil
}
