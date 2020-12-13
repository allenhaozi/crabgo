package mesh

import (
	crabcorev1 "github.com/allenhaozi/crabgo/pkg/apis/core/v1"
	"github.com/allenhaozi/crabgo/pkg/crab/dao/models"
	"github.com/allenhaozi/crabgo/pkg/register"
	"github.com/pkg/errors"
)

type AppDocService struct {
	appDocDao *models.AppDocDAO
}

func NewAppDocService(cfg *register.Config) *AppDocService {
	ins := &AppDocService{
		appDocDao: models.NewAppDocDAO(cfg),
	}
	return ins
}

func (ads *AppDocService) GetAppDoc(ctx register.Context, appId string) (*crabcorev1.AppDoc, error) {
	doc := &crabcorev1.AppDoc{}

	docModel, err := ads.appDocDao.GetAppDocById(ctx, appId)

	if err != nil {
		return nil, errors.Wrap(err, "get model failure from db")
	}

	doc = docModel.ToAppDoc()

	return doc, nil
}
