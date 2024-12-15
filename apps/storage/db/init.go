package db

import (
	"github.com/sanzashi987/nino-work/apps/storage/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

func ConnectDB() {
	instance := db.ConnectDB()
	instance.AutoMigrate(&model.File{}, &model.Bucket{})
}
