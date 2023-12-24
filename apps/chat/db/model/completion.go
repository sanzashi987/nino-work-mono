package model

import "github.com/cza14h/nino-work/pkg/model"

type MessageModel struct {
	model.BaseModel
	Content string `gorm:"column:contet;type:char(255)"`
	ReplyTo int    `gorm:"reply_to; default:0"`
	Deleted bool   `gorm:"index:default:false;type:boolean"`
}

func (msg MessageModel) TableName() string {
	return "messages"
}

type CompletionModel struct {
	model.BaseModel

	Messages []MessageModel `gorm:""`
	Deleted  bool           `gorm:"type:boolean;index;"`
	Count    int            `gorm:"column:count"`
}

func (com CompletionModel) TableName() string {
	return "completions"
}
