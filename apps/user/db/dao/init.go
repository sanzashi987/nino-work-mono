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
	conf := config.GetConfig()
	db, err := gorm.Open(sqlite.Open(conf.System.DbName), &gorm.Config{
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

func newDBSession(ctx context.Context) *gorm.DB {
	return instance.WithContext(ctx)
}
