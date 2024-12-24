package dao

import (
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"gorm.io/gorm"
)

func ConnectDB(name ...string) {
	instance := db.ConnectDB(name...)
	migrateTable(instance)
}

func migrateTable(db *gorm.DB) {
	db.AutoMigrate(&model.UserModel{}, &model.PermissionModel{}, &model.RoleModel{}, &model.ApplicationModel{})
	defaultRecord(db)
}

func defaultRecord(db *gorm.DB) {

	var roles, apps int64
	permission := model.PermissionModel{
		Name: "Root SuperAdmin",
		Code: "root.super_admin",
	}
	db.Model(&model.RoleModel{}).Count(&roles)
	if roles == 0 {

		// Create default user
		user := model.UserModel{
			Username: "admin",
			Password: "admin",
		}

		// Create default role
		role := &model.RoleModel{
			Name:        "Root SuperAdmin",
			Code:        "root.super_admin",
			Permissions: []model.PermissionModel{permission},
			Users:       []model.UserModel{user},
		}

		db.Create(role)
	}

	db.Model(&model.ApplicationModel{}).Count(&apps)
	if apps == 0 {
		// Create default application
		application := &model.ApplicationModel{
			Name:        "Root",
			Code:        "root",
			Description: "Root application",
			Status:      model.SystemOnline,
			CreateBy:    1,
			SuperAdmin:  permission.Id,
			Admin:       permission.Id,
			Permissions: []model.PermissionModel{permission},
		}
		db.Create(application)
	}

}
