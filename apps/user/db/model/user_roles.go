package model

import "time"

type UserRoleModel struct {
	RoleId   uint64    `gorm:"column:role_id;index"`
	UserId   uint64    `gorm:"column:user_id;index"`
	ExpireAt time.Time `gorm:"column:expire_at"`
}

func (u UserRoleModel) TableName() string {
	return "user_roles"
}
