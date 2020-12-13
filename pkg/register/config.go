package register

import (
	"github.com/allenhaozi/crabgo/pkg/register/config"
	"github.com/allenhaozi/crabgo/pkg/register/storage"
)

var PConfig *Config

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

	err := c.initConfig()

	if err != nil {
		return nil, err
	}

	// set a global var save Config value
	PConfig = c

	return c, nil
}

func (c *Config) initConfig() error {
	c.GeneralConfig = &GeneralConfig{}

	cfg, err := config.GetCrabConfig()
	if err != nil {
		return err
	}

	// sage config
	c.GeneralConfig.CrabConfig = cfg

	return nil
}

func (c *Config) initClient() error {

	c.ExtraConfig = &ExtraConfig{}

	// mysql client
	my, err := storage.NewMySqlClient(c.GeneralConfig.MyConfig)
	if err != nil {
		return err
	}
	c.ExtraConfig.MyClient = my

	return nil
}
