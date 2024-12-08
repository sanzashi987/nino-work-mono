package dao

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/chat/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"gorm.io/gorm"
)

func ConnectDB() {
	instance := db.ConnectDB()
	migrateTable(instance)
}

func migrateTable(db *gorm.DB) {
	db.AutoMigrate(&model.DialogModel{}, &model.MessageModel{}, &model.UserConfigModel{})
}

func newDBSession(ctx context.Context) *gorm.DB {
	return db.NewDBSession(ctx)
}
