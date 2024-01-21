package model

import "github.com/cza14h/nino-work/pkg/db"

type ThemeModel struct {
	db.BaseModel
	Name      string `gorm:"column:name"`
	Deleted   int8   `gorm:"column:deleted"`
	Flag      int8   `gorm:"column:flag"` //0 is built-in
	Workspace uint64 `gorm:"column:workspace;index"`
	Config    string `gorm:"type:blob"`
}

func (theme ThemeModel) TableName() string {
	return "themes"
}
