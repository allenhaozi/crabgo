package storage

import (
	"github.com/allenhaozi/crabgo/pkg/register/config"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySqlClient struct {
	*gorm.DB
}

func NewMySqlClient(cfg *config.CrabMysqlConfig) (*MySqlClient, error) {
	instance := &MySqlClient{}
	db, err := gorm.Open(mysql.Open(cfg.GetConnString()), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	instance.DB = db

	return instance, nil
}
