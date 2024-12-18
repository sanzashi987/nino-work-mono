package dao

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type UserDao struct {
	db.BaseDao[model.UserModel]
}

func NewUserDao(ctx context.Context, dao ...*db.BaseDao[model.UserModel]) *UserDao {
	return &UserDao{BaseDao: db.NewDao[model.UserModel](ctx, dao...)}
}

func (dao *UserDao) FindUserById(id uint64) (user *model.UserModel, err error) {
	err = dao.GetOrm().Where("id = ?", id).First(user).Error
	return

}

func (dao *UserDao) FindUserByUsername(username string) (user *model.UserModel, err error) {
	err = dao.GetOrm().Where("username = ?", username).First(user).Error
	return
}

func (dao *UserDao) CreateUser(newUser *model.UserModel) (err error) {
	err = dao.GetOrm().Create(newUser).Error
	return
}

func (dao *UserDao) UpdateUser(nextUser *model.UserModel) (err error) {
	err = dao.GetOrm().Updates(nextUser).Error
	return
}

func (dao *UserDao) FindAllAdmins(id uint64) (*model.UserModel, error) {
	user := model.UserModel{}
	user.Id = id

	if err := dao.GetOrm().Model(&user).Association("Roles").Find(&user.Roles); err != nil {
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
