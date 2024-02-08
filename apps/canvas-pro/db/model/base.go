package model

import "github.com/cza14h/nino-work/pkg/db"

type BaseModel struct {
	db.BaseModel
	Name      string
	Workspace string
	Creator   string
	Deleted   uint8 `gorm:"deleted:tinyint(8)"`
}
