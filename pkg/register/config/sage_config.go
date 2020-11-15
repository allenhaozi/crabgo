package config

import (
	"os"
	"strings"
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

type Sage struct {
	SageNamespace          string
	CommonConfigFileServer string
}

type SageConfig struct {
	Sage
	SageGateWay               string
	SageRegistryAddr          string
	ConfigCenterAddr          string
	KeyStoneAddr              string
	SageLicenseAddr           string
	WebServerPort             string
	StorageBasePath           string
	LogLevel                  log.Level
	PeriodAppStatusInspection time.Duration
	DockerSecret              string
	RSAKeyDir                 string
}

func GetSageConfig() (*SageConfig, error) {

	cfg := &SageConfig{}
	err := cfg.initByEnv()
	if err != nil {
		return nil, errors.Wrap(err, "init config by env failed")
	}
	return cfg, nil
}

func (cfg *SageConfig) initByEnv() error {
	if gw, ok := os.LookupEnv("SAGE_GATEWAY"); ok {
		// TODO regular ?
		if !strings.HasPrefix(gw, httpPrefix) {
			gw = httpPrefix + gw
		}
		cfg.SageGateWay = gw
	} else {
		cfg.SageGateWay = defaultGateWay
	}


	return nil
}
