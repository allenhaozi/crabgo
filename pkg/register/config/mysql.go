package config

import (
	"errors"
	"fmt"
	"net/url"
	"os"
)

var (
	timeZone = "Asia/Shanghai"
	dbName   = "app-manager"
)

type CrabMysqlConfig struct {
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

func GetMysqlConfig() (*CrabMysqlConfig, error) {

	var cfg CrabMysqlConfig
	if h, ok := os.LookupEnv("DB_HOST"); ok {
		cfg.Host = h
	} else {
		return nil, errors.New("can not found db host")
	}
	if p, ok := os.LookupEnv("DB_PORT"); ok {
		cfg.Port = p
	} else {
		return nil, errors.New("can not found db port")
	}
	if pwd, ok := os.LookupEnv("DB_PWD"); ok {
		cfg.PassWord = pwd
	} else {
		return nil, errors.New("can not found db password")
	}
	if u, ok := os.LookupEnv("DB_USER"); ok {
		cfg.User = u
	} else {
		return nil, errors.New("can not found db user")
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

func (smc *CrabMysqlConfig) GetConnString() string {
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
