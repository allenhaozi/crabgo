package models

import (
	"time"

	crabcorev1 "github.com/allenhaozi/crabgo/pkg/apis/core/v1"
)

type AppDocModel struct {
	ID          uint32 `gorm:"AUTO_INCREMENT"`
	AppId       string
	Name        string
	Resource    string
	Namespace   string
	Labels      string
	LabelIndex  string
	Annotations string
	CTime       time.Time `gorm:"column:ctime"`
	MTime       time.Time `gorm:"column:mtime"`
}

func (ad AppDocModel) TableName() string {
	return "app_docs"
}

func (ad *AppDocModel) ToAppDoc() *crabcorev1.AppDoc {
	doc := &crabcorev1.AppDoc{}
	return doc
}
