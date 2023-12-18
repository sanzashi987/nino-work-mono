package dao

import (
	"context"

	"github.com/cza14h/nino-work/apps/user/db/model"
	"github.com/cza14h/nino-work/config"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var instance *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open(config.DbName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Fail to connect database")
	}
	instance = db
	migrateTable(db)
}

func migrateTable(db *gorm.DB) {
	db.AutoMigrate(&model.UserModel{})
}

func NewDBSession(ctx context.Context) *gorm.DB {
	return instance.WithContext(ctx)
}
