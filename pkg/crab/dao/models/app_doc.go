package models

import (
	"time"

	"github.com/allenhaozi/crabgo/pkg/register"
	"gorm.io/gorm"
)

type AppDocDAO struct {
	AbstractModel
}

func NewAppDocDAO(cfg *register.Config) *AppDocDAO {
	add := &AppDocDAO{}
	add.SetConf(cfg)
	return add
}

func (ad *AppDocDAO) GetAppDocById(ctx register.Context, appId string) (*AppDocModel, error) {
	var appDoc AppDocModel
	db := ad.getDB(ctx.DebugModel())
	resp := db.Where("app_id = ?", appId).First(&appDoc)
	return &appDoc, resp.Error
}

func (ad *AppDocDAO) Create(ctx register.Context, appDocModel *AppDocModel) *gorm.DB {
	db := ad.getDB(ctx.DebugModel())
	t := time.Now()
	appDocModel.CTime, appDocModel.MTime = t, t
	res := db.Create(appDocModel)
	return res
}
