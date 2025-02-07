package model

import "github.com/sanzashi987/nino-work/pkg/db"

const (
	NamespaceOwner  = 1 // 所有者
	NamespaceAdmin  = 2 // 管理员
	NamespaceMember = 3 // 普通成员
)

type NamespaceUserRelation struct {
	NamespaceID  uint64 `gorm:"primaryKey;column:namespace_id"`
	UserID       uint64 `gorm:"primaryKey;column:user_id"`
	RelationType int    `gorm:"column:relation_type"` // 1:所有者 2:管理员 3:普通成员
}

func (r NamespaceUserRelation) TableName() string {
	return "namespace_users"
}

type NamespaceModel struct {
	db.BaseModel
	Name        string   `gorm:"column:name;type:varchar(255)"`
	Code        string   `gorm:"column:code;type:varchar(255);uniqueIndex"`
	Description string   `gorm:"column:description"`
	Status      int      `gorm:"column:status;default:0"` // 0:启用 1:禁用
	Users       []uint64 `gorm:"-"`
}

func (n NamespaceModel) TableName() string {
	return "config_namespaces"
}
