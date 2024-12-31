package dao

import (
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"github.com/sanzashi987/nino-work/pkg/utils"
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

	var roles, apps, permissions int64

	tx := db.Begin()

	tx.Model(&model.RoleModel{}).Count(&roles)
	tx.Model(&model.ApplicationModel{}).Count(&apps)
	tx.Model(&model.PermissionModel{}).Count(&permissions)
	if roles == 0 && apps == 0 && permissions == 0 {

		// Create default user
		user := model.UserModel{
			Username: "admin",
			Password: "admin",
		}

		adminRole, rootPermission := model.CreateRoleWithPermission("Root Super Admin", "root.admin.super")

		codes := []string{"user", "app", "role", "permission"}
		userRoles := []model.RoleModel{adminRole}
		for _, code := range codes {
			role, _ := model.CreateRoleWithPermission(
				utils.Capitialize(code)+" Admin Role",
				"root.admin."+code,
			)
			userRoles = append(userRoles, role)
		}

		user.Roles = userRoles
		tx.Create(&user)

		menus := []*model.MenuModel{}
		for _, code := range codes {
			menu := &model.MenuModel{
				Name:   utils.Capitialize(code) + " Management",
				Code:   "root.management." + code,
				Type:   model.MenuTypeMenu,
				Order:  0,
				Status: model.MenuEnable,
				Path:   "/dashboard/manage/" + code,
			}
			menus = append(menus, menu)
		}

		tx.Create(&menus)
		for index := range codes {
			menu := menus[index]
			tx.Model(menu).Association("Roles").Append(&user.Roles[index+1])
		}

		// Create default application
		application := &model.ApplicationModel{
			Name:        "Root",
			Code:        "root.nino.work",
			Description: "Root application",
			Status:      model.SystemOnline,
			CreateBy:    user.Id,
			SuperAdmin:  rootPermission.Id,
			Admin:       rootPermission.Id,
		}
		tx.Create(application)
		toUpdate := map[string]any{"app_id": application.Id}
		tx.Model(rootPermission).Updates(toUpdate)
	}

	tx.Commit()

}
