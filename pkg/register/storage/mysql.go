package storage

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/allenhaozi/crabgo/pkg/register/config"
)

type MySqlClient struct {
	*gorm.DB
}

func NewMySqlClient(cfg *config.SageMysqlConfig) (*MySqlClient, error) {
	instance := &MySqlClient{}
	db, err := gorm.Open("mysql", cfg.GetConnString())

	if err != nil {
		return nil, err
	}
	instance.DB = db

	return instance, nil
}
