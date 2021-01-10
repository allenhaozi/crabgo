package models

import (
	"github.com/allenhaozi/crabgo/pkg/register"
	"github.com/allenhaozi/crabgo/pkg/register/storage"
	"gorm.io/gorm"
)

type AbstractModel struct {
	Cfg *register.Config
	db  *storage.MySqlClient
}

func (am *AbstractModel) SetConf(cfg *register.Config) {
	if cfg != nil {
		am.Cfg = cfg
	}
	if cfg.ExtraConfig.MyClient != nil {
		am.db = cfg.ExtraConfig.MyClient
	}
}

func (am *AbstractModel) getDB(debug bool) *gorm.DB {
	db := am.db.DB
	if debug {
		db = db.Debug()
	}
	return db
}
