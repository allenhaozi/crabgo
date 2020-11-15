package models

/*
import (
	"github.com/jinzhu/gorm"
	"githu.com/allenhaozi/crabgo/pkg/register"
	"github.com/allenhaozi/pkg/register/storage"
)

type AbstractModel struct {
	Cfg *register.Config
	db  *storage.MySqlClient
}

func (ma *AbstractModel) SetConf(cfg *register.Config) {
	ma.Cfg = cfg
	ma.db = cfg.ExtraConfig.MyClient
}

func (ma *AbstractModel) getDB(debug bool) *gorm.DB {
	db := ma.db.DB
	if debug {
		db = db.Debug()
	}
	return db
}
*/
