package db

import (
	"context"

	"github.com/cza14h/nino-work/config"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var instance *gorm.DB

func ConnectDB(names ...string) *gorm.DB {
	fallbackName := ""
	if len(names) >= 1 {
		fallbackName = names[0]
	} else {
		conf := config.GetConfig()
		fallbackName = conf.System.DbName
	}

	db, err := gorm.Open(sqlite.Open(fallbackName+".db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Fail to connect database")
	}
	instance = db
	return db
}

func NewDBSession(ctx context.Context) *gorm.DB {
	return instance.WithContext(ctx)
}
