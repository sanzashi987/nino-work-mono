package dao

import (
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"gorm.io/gorm"
)

func FindUserById(tx *gorm.DB, id uint64) (*model.UserModel, error) {
	user := &model.UserModel{}

	err := tx.Model(user).Where("id = ?", id).First(user).Error
	return user, err

}

func FindUserByUsername(tx *gorm.DB, username string) (*model.UserModel, error) {
	user := &model.UserModel{}
	err := tx.Model(&model.UserModel{}).Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(tx *gorm.DB, newUser *model.UserModel) error {
	return tx.Create(newUser).Error
}

func UpdateUser(tx *gorm.DB, nextUser *model.UserModel) error {
	return tx.Updates(nextUser).Error

}

func FindUserWithRoles(tx *gorm.DB, id uint64) (*model.UserModel, error) {
	user := model.UserModel{}
	user.Id = id
	if err := tx.Model(&user).Find(&user).Error; err != nil {
		return nil, err
	}

	var userRoles []*model.UserRoleModel
	if err := tx.Model(&model.UserRoleModel{UserId: id}).Find(&userRoles).Error; err != nil {
		return nil, err
	}

	roleIds := make([]uint64, len(userRoles))
	for i, ur := range userRoles {
		roleIds[i] = ur.RoleId
	}
	var roles []*model.RoleModel
	if len(roleIds) > 0 {
		if err := tx.Model(&model.RoleModel{}).Where("id IN ?", roleIds).Find(&roles).Error; err != nil {
			return nil, err
		}
	}
	user.Roles = roles

	return &user, nil
}
