package models

import (
	"time"

	"github.com/jinzhu/copier"

	crabcorev1 "github.com/allenhaozi/crabgo/pkg/apis/core/v1"
)

type AppDocModel struct {
	ID          uint32 `gorm:"AUTO_INCREMENT"`
	Name        string
	AppId       string
	AppName     string
	AppVersion  string
	AppGroup    string
	User        string `gorm:"column:create_user"`
	UserId      uint32 `gorm:"column:create_user_id"`
	Namespace   string
	Labels      string
	Annotations string
	Status      int
	CTime       time.Time `gorm:"column:ctime"`
	MTime       time.Time `gorm:"column:mtime"`
}

func (ad AppDocModel) TableName() string {
	return "app_docs"
}

func (ad *AppDocModel) ToAppDoc() (*crabcorev1.AppDoc, error) {
	doc := &crabcorev1.AppDoc{}
	if err := copier.Copy(doc, ad); err != nil {
		return nil, err
	}

	return doc, nil
}

func (ad *AppDocModel) ToModel(doc *crabcorev1.AppDoc) error {
	if err := copier.Copy(ad, doc); err != nil {
		return err
	}

	return nil
}
