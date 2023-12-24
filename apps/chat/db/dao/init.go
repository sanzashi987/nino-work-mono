package dao

import (
	"context"

	"github.com/cza14h/nino-work/apps/chat/db/model"
	"github.com/cza14h/nino-work/pkg/db"
	"gorm.io/gorm"
)

func ConnectDB() {
	instance := db.ConnectDB()
	migrateTable(instance)
}

func migrateTable(db *gorm.DB) {
	db.AutoMigrate(&model.DialogModel{}, &model.MessageModel{})
}

func NewDBSession(ctx context.Context) *gorm.DB {
	return db.NewDBSession(ctx)
}
