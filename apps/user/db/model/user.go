package model

import (
	"github.com/cza14h/nino-work/pkg/db"
	"github.com/cza14h/nino-work/pkg/utils"
	"gorm.io/gorm"
)

const (
	User  = 0
	Admin = 1
)

type UserModel struct {
	db.BaseModel
	Username string `gorm:"column:username;type:varchar(255);unique"`
	Password string `gorm:"column:password;type:varchar(255)"`
	Role     int32  `gorm:"column:role"`
	Features string `gorm:"column:features;type:varchar(255)"`
	Fobidden bool   `gorm:"column:forbidden"`
}

func (u UserModel) TableName() string {
	return "users"
}

func (user *UserModel) CheckPassowrd(password string) bool {
	return utils.CompareHash(user.Password, password)
}

// Gorm hook
func (user *UserModel) BeforeSave(tx *gorm.DB) (err error) {
	if !utils.IsHashed(user.Password) {
		user.Password = utils.MakeHash(user.Password)
	}
	return
}
