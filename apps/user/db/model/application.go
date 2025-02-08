package model

import (
	"github.com/sanzashi987/nino-work/pkg/db"
	"github.com/sanzashi987/nino-work/pkg/utils"
	"gorm.io/gorm"
)

const (
	SystemOnline  = 0
	SystemOffline = 1
)

type ApplicationModel struct {
	db.BaseModel
	Name        string `gorm:"column:name;type:varchar(255)"`
	Code        string `gorm:"column:code;type:varchar(255);uniqueIndex"`
	Description string `gorm:"column:description"`
	Status      uint   `gorm:"column:status"`
	CreateBy    uint64 `gorm:"column:create_by"`
	AccessKey   string `gorm:"column:access_key;type:varchar(32);uniqueIndex"` // AK
	SecretKey   string `gorm:"column:secret_key;type:varchar(64)"`             // SK
	ServiceUser uint64 `gorm:"column:service_user"`
	// store permission Id here
	SuperAdmin uint64 `gorm:"column:super_admin;"`
	// store permission Id here
	Admin       uint64             `gorm:"column:admin;"`
	Permissions []*PermissionModel `gorm:"foreignKey:AppId"`
}

func (f ApplicationModel) TableName() string {
	return "applications"
}

func (app *ApplicationModel) BeforeCreate(tx *gorm.DB) error {
	// 生成 AK、SK
	accessKey, secretKey, err := utils.GenerateSecureKeys()

	if err != nil {
		return err
	}
	app.AccessKey, app.SecretKey = accessKey, secretKey

	// 创建应用用户
	appUser := &UserModel{
		Username: app.Code,
		Password: "",
		Type:     Application,
	}

	if err := tx.Create(appUser).Error; err != nil {
		return err
	}

	app.ServiceUser = appUser.Id
	return nil

}
