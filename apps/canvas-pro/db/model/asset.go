package model

import "github.com/cza14h/nino-work/pkg/db"

const FILE = 0
const FONT = 1

type AssetModel struct {
	db.BaseModel
	Deleted   uint8 `gorm:"deleted:tinyint(8)"`
	Name      string
	Workspace string
	Creator   string
	Type      uint8
	FileId    string
	FilePath  string
}

type AssetGroupModel struct {
}
