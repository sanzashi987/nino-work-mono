package dao

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{DB: newDBSession(ctx)}
}

func (dao *UserDao) FindUserById(id uint64) (user *model.UserModel, err error) {
	err = dao.Where("id = ?", id).First(user).Error
	return

}

func (dao *UserDao) FindUserByUsername(username string) (user *model.UserModel, err error) {
	err = dao.Where("username = ?", username).First(user).Error
	return
}

func (dao *UserDao) CreateUser(newUser *model.UserModel) (err error) {
	err = dao.Create(newUser).Error
	return
}

func (dao *UserDao) UpdateUser(nextUser *model.UserModel) (err error) {
	err = dao.Updates(nextUser).Error
	return
}

func (dao *UserDao) FindAllAdmins(id uint64) (*model.UserModel, error) {
	user := model.UserModel{}
	user.Id = id

	if err := dao.Model(&user).Association("Roles").Find(&user.Roles); err != nil {
		return nil, err
	}

	admins := []model.UserModel{}
	for _, role := range user.Roles {
		if role.Key == "admin" {
			admins = append(admins, user)
		}
	}

	

	return nil, nil
}
