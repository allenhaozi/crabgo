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

package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type LogConfig struct {
	LogLevel log.Level
}

func (l *CrabConfig) initLogConfig() error {
	cfg := &LogConfig{}

	// log level
	if l, ok := os.LookupEnv("LOG_LEVEL"); ok {
		level, err := log.ParseLevel(l)
		if err != nil {
			return err
		}
		cfg.LogLevel = level
	} else {
		cfg.LogLevel = log.InfoLevel
	}

	l.LogConfig = cfg

	return nil
}
