package db

import "github.com/cza14h/nino-work/pkg/model"

type UserModel struct {
	model.BaseModel
	Username string `gorm:"column:username;type:varchar(255);unique"`
	Password string `gorm:"column:password;type:varchar(255)"`
	Role     int32    `gorm:"column:role"`
}

func (u UserModel) TableName() string {
	return "users"
}
