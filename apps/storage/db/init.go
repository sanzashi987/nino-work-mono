package db

import (
	"github.com/sanzashi987/nino-work/apps/storage/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

func ConnectDB(name ...string) {
	instance := db.ConnectDB(name...)
	instance.AutoMigrate(&model.File{}, &model.Bucket{})
}
