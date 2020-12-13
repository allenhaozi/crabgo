package config

import (
	"os"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var (
	defaultGateWay                   = "http://gateway"
	httpPrefix                       = "http://"
	defaultWebPort                   = "9000"
	defaultPeriodAppStatusInspection = 10
	defConfigCenterAddr              = "http://config-center"
	defKeyStoneAddr                  = "http://keystone"
	defRSAKeyDir                     = "/home/work/etc"
)

type Crab struct {
	Namespace string
}

type CrabConfig struct {
	Crab
	CrabGateWay               string
	CrabRegistryAddr          string
	ConfigCenterAddr          string
	KeyStoneAddr              string
	CrabLicenseAddr           string
	WebServerPort             string
	StorageBasePath           string
	LogLevel                  log.Level
	PeriodAppStatusInspection time.Duration
	DockerSecret              string
	RSAKeyDir                 string
}

func GetCrabConfig() (*CrabConfig, error) {

	cfg := &CrabConfig{}
	err := cfg.initByEnv()
	if err != nil {
		return nil, errors.Wrap(err, "init config by env failed")
	}
	return cfg, nil
}

func (cfg *CrabConfig) initByEnv() error {
	if port, ok := os.LookupEnv("WEB_SERVER_PORT"); ok {
		cfg.WebServerPort = port
	} else {
		cfg.CrabGateWay = defaultGateWay
	}

	return nil
}
