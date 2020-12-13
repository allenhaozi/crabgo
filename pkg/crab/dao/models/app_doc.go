package models

import "github.com/allenhaozi/crabgo/pkg/register"

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
