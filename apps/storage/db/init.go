package db

import "github.com/sanzashi987/nino-work/pkg/db"

func ConnectDB() {
	instance := db.ConnectDB()
	instance.AutoMigrate(&File{})
}
