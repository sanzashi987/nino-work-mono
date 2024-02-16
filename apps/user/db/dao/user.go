package dao

import (
	"context"

	"github.com/cza14h/nino-work/apps/user/db/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{DB: newDBSession(ctx)}
}

func (dao *UserDao) FindUserById(id string) (user *model.UserModel, err error) {
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

