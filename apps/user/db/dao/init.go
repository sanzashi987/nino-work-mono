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

	var roleCounts, appCounts, permissionCounts int64

	tx := db.Begin()

	tx.Model(&model.RoleModel{}).Count(&roleCounts)
	tx.Model(&model.ApplicationModel{}).Count(&appCounts)
	tx.Model(&model.PermissionModel{}).Count(&permissionCounts)
	if roleCounts == 0 && appCounts == 0 && permissionCounts == 0 {

		// Create default user
		user := model.UserModel{
			Username: "admin",
			Password: "admin",
		}

		permissionsToCreate := []*model.PermissionModel{}
		adminRole, rootPermission := model.CreateRoleWithPermission("Root Super Admin", "root.admin.super")
		// permissionsToCreate = append(permissionsToCreate, &rootPermission)

		codes := []string{"user", "app", "role"}
		userRoles := []*model.RoleModel{adminRole}
		for _, code := range codes {
			role, permission := model.CreateRoleWithPermission(
				utils.Capitialize(code)+" Admin Role",
				"root.admin."+code,
			)
			permissionsToCreate = append(permissionsToCreate, permission)
			userRoles = append(userRoles, role)
		}

		user.Roles = userRoles
		if err := tx.Create(&user).Error; err != nil {
			tx.Rollback()
			return
		}

		menus := []*model.MenuModel{}
		for _, code := range codes {
			menu := &model.MenuModel{
				Name:   utils.Capitialize(code) + " Management",
				Code:   "root.management." + code,
				Type:   model.MenuTypeMenu,
				Order:  0,
				Status: model.MenuEnable,
				Path:   "/home/root/" + code,
			}
			menus = append(menus, menu)
		}

		if err := tx.Create(&menus).Error; err != nil {
			tx.Rollback()
			return
		}
		for index := range codes {
			menu := menus[index]
			if err := tx.Model(menu).Association("Roles").Append(&user.Roles[index+1]); err != nil {
				tx.Rollback()
				return
			}
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

		if err := tx.Create(application).Error; err != nil {
			tx.Rollback()
			return
		}
		if err := tx.Model(application).Association("Permissions").Append(&permissionsToCreate); err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()

}
