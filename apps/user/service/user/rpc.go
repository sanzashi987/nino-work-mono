package userService

import (
	"context"
	"errors"
	"time"

	"github.com/sanzashi987/nino-work/apps/user/consts"
	"github.com/sanzashi987/nino-work/apps/user/db/dao"
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/controller"
	"github.com/sanzashi987/nino-work/pkg/db"
	"github.com/sanzashi987/nino-work/proto/user"
	"gorm.io/gorm"
)

type UserServiceRpc struct {
	user.UserServiceHandler
}

var UserServiceRpcImpl = &UserServiceRpc{}

// func (u *UserServiceRpc) GetApplicationPermissions() {

// }
// func (u *UserServiceRpc) GetUserPermissions() {
// }

func (u *UserServiceRpc) UserLogin(ctx context.Context, in *user.UserLoginRequest, out *user.UserLoginResponse) (err error) {
	user, err := dao.FindUserByUsername(db.NewTx(ctx), in.Username)
	if err != nil {
		out.Reason = consts.UsernameNotExist
		return
	}

	if valid := user.CheckPassowrd(in.Password); !valid {
		err = errors.New("user password not match")
		out.Reason = consts.PasswordNotMatch
		return
	}

	var token string

	var days int
	if in.Expiry != nil {
		days = int(*in.Expiry)
	} else {
		days = 1
	}

	token, err = controller.GenerateToken(user.Username, user.Id, time.Hour*24*time.Duration(days))
	if err != nil {
		out.Reason = consts.FailToCreateToken
		return
	}

	out.Expiry = int32(days)
	out.JwtToken = token
	out.Reason = consts.Success
	return
}

func (u *UserServiceRpc) UserRegister(ctx context.Context, in *user.UserRegisterRequest, out *user.UserLoginResponse) (err error) {
	tx := db.NewTx(ctx)

	user, err := dao.FindUserByUsername(tx, in.Username)
	if user != nil {
		out.Reason = consts.UsernameExisted
		err = errors.New("username existed")
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {

		if in.Password != in.PasswordConfirm {
			out.Reason = consts.PasswordNotMatch
			err = errors.New("password does not match")
			return
		}

		user := model.UserModel{
			Username: in.Username,
			Password: in.Password,
			Fobidden: false,
		}

		dao.CreateUser(tx, &user)
		var token string
		token, err = controller.GenerateToken(user.Username, user.Id)
		if err != nil {
			out.Reason = consts.FailToCreateToken
			return
		}
		out.JwtToken = token
		return
	}

	out.Reason = consts.InternalServerError
	err = errors.New("unknown edge case in user service")
	return
}

// GetApplicationPermissions implements user.UserServiceHandler.
func (u *UserServiceRpc) GetApplicationPermissions(context.Context, *user.ApplicationPermissionsRequest, *user.ApplicationPermissionsResponse) error {
	panic("unimplemented")
}

// GetUserPermissions implements user.UserServiceHandler.
func (u *UserServiceRpc) GetUserPermissions(context.Context, *user.UserPermissionsRequest, *user.UserPermissionsResponse) error {
	panic("unimplemented")
}
