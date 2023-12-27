package model

import "github.com/cza14h/nino-work/pkg/db"

type MessageModel struct {
	db.BaseModel
	DialogID uint64 `gorm:"index"`
	Content  string `gorm:"column:contet;type:char(255)"`
	ReplyTo  uint64 `gorm:"column:reply_to;default:0"`
	Deleted  bool   `gorm:"index:default:false;type:boolean"`
}

func (msg MessageModel) TableName() string {
	return "messages"
}

type DialogModel struct {
	db.BaseModel
	Messages []MessageModel `gorm:"foreignKey:DialogID;"`
	Deleted  bool           `gorm:"type:boolean;index;"`
	Count    int            `gorm:"column:count"`
	/**
	json string for dto.DialogPreference
	*/
	Preference string `gorm:"column:preference;type:char(255)"`
}

func (com DialogModel) TableName() string {
	return "dialogs"
}

type UserConfigModel struct {
	UserId uint64 `gorm:"index;primaryKey"`
	/**
	json string for dto.UserPreference
	*/
	Preference string `gorm:"column:preference;type:char(255)"`
}

func (user UserConfigModel) TableName() string {
	return "user_config"
}
