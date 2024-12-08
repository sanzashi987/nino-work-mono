package dao

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"gorm.io/gorm"
)


func ConnectDB() {
	instance := db.ConnectDB()
	migrateTable(instance)
}

func migrateTable(db *gorm.DB) {
	db.AutoMigrate(&model.UserModel{})
}

func newDBSession(ctx context.Context) *gorm.DB {
	return db.NewDBSession(ctx)
}
