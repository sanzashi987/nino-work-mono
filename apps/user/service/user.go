package service

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/sanzashi987/nino-work/apps/user/db/dao"
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/auth"
	"github.com/sanzashi987/nino-work/pkg/db"
	"github.com/sanzashi987/nino-work/proto/user"
	"gorm.io/gorm"
)

type UserServiceRpc struct{}

var UserServiceRpcImpl *UserServiceRpc = &UserServiceRpc{}
var once sync.Once

func GetUserServiceRpc() user.UserServiceHandler {
	return UserServiceRpcImpl
}

func (u *UserServiceRpc) UserLogin(ctx context.Context, in *user.UserLoginRequest, out *user.UserLoginResponse) (err error) {
	user, err := dao.NewUserDao(ctx).FindUserByUsername(in.Username)
	if err != nil {
		out.Reason = UsernameNotExist
		return
	}

	if valid := user.CheckPassowrd(in.Password); !valid {
		out.Reason = PasswordNotMatch
		return
	}

	var token string

	var days int
	if in.Expiry != nil {
		days = int(*in.Expiry)
	} else {
		days = 1
	}

	token, err = auth.GenerateToken(user.Username, user.Id, time.Hour*24*time.Duration(days))
	if err != nil {
		out.Reason = FailToCreateToken
		return
	}

	out.Expiry = int32(days)
	out.JwtToken = token
	out.Reason = Success
	return
}

func (u *UserServiceRpc) UserRegister(ctx context.Context, in *user.UserRegisterRequest, out *user.UserLoginResponse) (err error) {
	dbSession := dao.NewUserDao(ctx)
	user, err := dbSession.FindUserByUsername(in.Username)
	if user != nil {
		out.Reason = UsernameExisted
		err = errors.New("Username existed")
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {

		if in.Password != in.PasswordConfirm {
			out.Reason = PasswordNotMatch
			err = errors.New("Password does not match")
			return
		}

		user := model.UserModel{
			Username: in.Username,
			Password: in.Password,
			Fobidden: false,
		}

		dbSession.CreateUser(&user)
		var token string
		token, err = auth.GenerateToken(user.Username, user.Id)
		if err != nil {
			out.Reason = FailToCreateToken
			return
		}
		out.JwtToken = token
		return
	}

	out.Reason = InternalServerError
	err = errors.New("Unknown edge case in user service")
	return
}

type UserServiceWeb struct{}

var UserServiceWebImpl *UserServiceWeb = &UserServiceWeb{}

type UserInfoResponse struct {
	UserId   uint64 `json:"user_id"`
	Username string `json:"username"`
	Features string `json:"features"`
}

func (u *UserServiceWeb) UserInfo(ctx context.Context, userId uint64) (*UserInfoResponse, error) {
	if userId == 0 {
		return nil, errors.New("user id is equired")
	}

	user, err := dao.NewUserDao(ctx).FindUserById(userId)
	if err != nil {
		return nil, err
	}

	return &UserInfoResponse{
		UserId:   user.Id,
		Username: user.Username,
	}, nil
}

func (u *UserServiceWeb) GetUserRoles(ctx context.Context, userId uint64) ([]model.RoleModel, error) {

	userDao := dao.NewUserDao(ctx)

	user, err := userDao.FindUserWithRoles(userId)
	if err != nil {
		return nil, err
	}

	return user.Roles, nil
}

func (u *UserServiceWeb) GetUserRoleWithPermissions(ctx context.Context, userId uint64) ([]model.RoleModel, error) {
	roles, _, err := getUserRolePermission(ctx, userId)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func getUserRolePermission(ctx context.Context, userId uint64) ([]model.RoleModel, *db.BaseDao[model.RoleModel], error) {
	userDao := dao.NewUserDao(ctx)
	roleDao := dao.NewRoleDao(ctx, (*db.BaseDao[model.RoleModel])(&userDao.BaseDao))
	user, err := userDao.FindUserWithRoles(userId)
	if err != nil {
		return nil, nil, err
	}

	userRoles := []*model.RoleModel{}
	for _, role := range user.Roles {
		userRoles = append(userRoles, &role)
	}

	err = roleDao.FindRolesWithPermissions(userRoles...)
	if err != nil {
		return nil, nil, err
	}

	return user.Roles, &roleDao.BaseDao, nil
}

func getUserAdmins(ctx context.Context, userId uint64) (*[]model.ApplicationModel, *db.BaseDao[model.RoleModel], error) {
	roles, roleDao, err := getUserRolePermission(ctx, userId)
	if err != nil {
		return nil, nil, err
	}

	applications := map[uint64]bool{}
	permissions := map[uint64]bool{}
	for _, role := range roles {
		for _, permission := range role.Permissions {
			applications[permission.AppId] = true
			permissions[permission.Id] = true
		}
	}

	appIds := []uint64{}
	for appId := range applications {
		appIds = append(appIds, appId)
	}
	apps := []model.ApplicationModel{}
	err = roleDao.GetOrm().Table("applications").Where("id IN ?", appIds).Find(&apps).Error
	if err != nil {
		return nil, nil, err

	}

	res := []model.ApplicationModel{}

	for _, app := range apps {
		_, superExist := permissions[app.SuperAdmin]
		_, adminExist := permissions[app.Admin]
		if superExist || adminExist {
			res = append(res, app)
		}
	}

	return &res, roleDao, nil
}
