package dao

import (
	"errors"

	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"gorm.io/gorm"
)

func CreateApp(tx *gorm.DB, app *model.ApplicationModel) error {
	// 检查是否存在相同Code的系统
	var existingSystem model.ApplicationModel
	err := tx.Where("code = ?", app.Code).First(&existingSystem).Error
	if err == nil {
		return errors.New("system code already exists")
	}

	return tx.Create(app).Error
}

func InitPermissionForApp(tx *gorm.DB, app *model.ApplicationModel, super *model.PermissionModel, admin *model.PermissionModel) error {

	toUpdate := map[string]any{
		"super_admin": super.Id,
		"admin":       admin.Id,
	}
	err := tx.Model(app).Updates(toUpdate).Error
	return err
}

func FindApplicationByIdWithPermission(tx *gorm.DB, id uint64) (*model.ApplicationModel, error) {
	app := model.ApplicationModel{}
	err := tx.Preload("Permissions").Where("id = ?", id).First(&app).Error
	return &app, err

}
