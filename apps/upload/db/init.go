package db

import "github.com/cza14h/nino-work/pkg/db"

func ConnectDB() {
	instance := db.ConnectDB()
	instance.AutoMigrate(&File{})
}
