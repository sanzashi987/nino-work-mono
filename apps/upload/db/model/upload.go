package model

import "github.com/cza14h/nino-work/pkg/db"

type File struct {
	db.BaseModel
	URL string `gorm:"type:varchar(255);unique;index"`
}
