package dao

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/pkg/db"
	"gorm.io/gorm"
)

func ConnectDB(dbname string) {
	instance := db.ConnectDB(dbname)
	instance.AutoMigrate(&model.ThemeModel{}, &model.WorkSpaceModel{})
}

func newDBSession(ctx context.Context) *gorm.DB {
	return db.NewDBSession(ctx)
}
