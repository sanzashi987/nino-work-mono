package model

import "github.com/cza14h/nino-work/pkg/db"

const DESGIN = 1
const FILE = 2
const FONT = 3
const BLOCK = 4
const COMPONENT = 5

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
	Type      uint8
	Deleted   uint8 `gorm:"deleted:tinyint(8)"`
	Workspace string
}
