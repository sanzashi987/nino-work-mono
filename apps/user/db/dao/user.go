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
	if err := tx.Model(&user).Preload("Roles").Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
