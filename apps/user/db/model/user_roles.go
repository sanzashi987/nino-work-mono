package model

import "time"

type UserRoleModel struct {
	Id       uint64 `gorm:"primaryKey"`
	RoleId   uint64 `gorm:"column:role_id;index"`
	UserId   uint64 `gorm:"column:user_id;index"`
	ExpireAt time.Time
}

func (u UserRoleModel) TableName() string {
	return "user_roles"
}
