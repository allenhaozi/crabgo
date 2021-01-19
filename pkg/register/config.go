package register

import (
	"github.com/allenhaozi/crabgo/pkg/register/config"
	"github.com/allenhaozi/crabgo/pkg/register/storage"
)

var pConfig *Config

func GetGlobalConfig() *Config {
	return pConfig
}

type ExtraConfig struct {
	MyClient *storage.MySqlClient
}

type GeneralConfig struct {
	CrabConfig *config.CrabConfig
	MyConfig   *config.CrabMysqlConfig
}

type Config struct {
	ExtraConfig   *ExtraConfig
	GeneralConfig *GeneralConfig
}

func NewConfig() (*Config, error) {

	c := &Config{}

	if err := c.initConfig(); err != nil {
		return nil, err
	}

	if err := c.initClient(); err != nil {
		return nil, err
	}

	// set a global var save Config value
	pConfig = c

	return c, nil
}

func (c *Config) initConfig() error {
	c.GeneralConfig = &GeneralConfig{}

	cfg, err := config.GetCrabConfig()
	if err != nil {
		return err
	}

	mCfg, err := config.GetMysqlConfig()
	if err != nil {
		return err
	}

	// assign mysql config
	c.GeneralConfig.MyConfig = mCfg

	// assign sys config
	c.GeneralConfig.CrabConfig = cfg

	return nil
}

func (c *Config) initClient() error {

	c.ExtraConfig = &ExtraConfig{}

	// mysql client

	if c.GeneralConfig.MyConfig != nil {
		my, err := storage.NewMySqlClient(c.GeneralConfig.MyConfig)
		if err != nil {
			return err
		}
		c.ExtraConfig.MyClient = my

	}

	return nil
}
