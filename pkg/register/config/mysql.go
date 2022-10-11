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
	"errors"
	"fmt"
	"net/url"
	"os"
)

var (
	timeZone = "Asia/Shanghai"
	dbName   = "openaios-iam"
)

type MysqlConfig struct {
	User       string
	Host       string
	PassWord   string
	Port       string
	DBName     string
	ConnParams string
}

// example : https://pkg.go.dev/github.com/jinzhu/gorm?tab=doc#Open
// db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
var connFormat = "%s:%s@tcp(%s:%s)/%s?%s"
var connParamsFormat = "charset=utf8&parseTime=True&loc=%s"

func GetMysqlConfig() (*MysqlConfig, error) {
	var cfg MysqlConfig
	if h, ok := os.LookupEnv("DB_HOST"); ok {
		cfg.Host = h
	} else {
		return nil, errors.New("can not find db host")
	}
	if p, ok := os.LookupEnv("DB_PORT"); ok {
		cfg.Port = p
	} else {
		return nil, errors.New("can not find db port")
	}
	if pwd, ok := os.LookupEnv("DB_PWD"); ok {
		cfg.PassWord = pwd
	} else {
		return nil, errors.New("can not find db password")
	}
	if u, ok := os.LookupEnv("DB_USER"); ok {
		cfg.User = u
	} else {
		return nil, errors.New("can not find db user")
	}

	if params, ok := os.LookupEnv("GO_MYSQL_PARAMS"); ok {
		cfg.ConnParams = params
	} else {
		var tz string
		if envTZ, ok := os.LookupEnv("TZ"); ok {
			tz = envTZ
		} else {
			// default tz
			tz = timeZone
		}
		cfg.ConnParams = fmt.Sprintf(connParamsFormat, url.QueryEscape(tz))
	}

	cfg.DBName = dbName

	return &cfg, nil
}

func (smc *MysqlConfig) GetConnString() string {
	return fmt.Sprintf(
		connFormat,
		smc.User,
		smc.PassWord,
		smc.Host,
		smc.Port,
		smc.DBName,
		smc.ConnParams,
	)
}
