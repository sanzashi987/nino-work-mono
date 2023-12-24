package model

import "github.com/cza14h/nino-work/pkg/db"

type MessageModel struct {
	db.BaseModel
	DialogID uint64 `gorm:"index"`
	Content string `gorm:"column:contet;type:char(255)"`
	ReplyTo int    `gorm:"column:reply_to;default:0"`
	Deleted bool   `gorm:"index:default:false;type:boolean"`
}

func (msg MessageModel) TableName() string {
	return "messages"
}

type DialogModel struct {
	db.BaseModel
	Messages []MessageModel `gorm:"foreignKey:DialogID;"`
	Deleted  bool           `gorm:"type:boolean;index;"`
	Count    int            `gorm:"column:count"`
}

func (com DialogModel) TableName() string {
	return "dialogs"
}
