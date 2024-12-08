package service

import (
	"context"
	"errors"
	"sync"

	"github.com/sanzashi987/nino-work/apps/user/db/dao"
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/auth"
	"github.com/sanzashi987/nino-work/proto/user"
	"gorm.io/gorm"
)

type UserServiceRpc struct{}

var UserServiceRpcImpl *UserServiceRpc
var once sync.Once

func GetUserServiceRpc() user.UserServiceHandler {
	once.Do(func() {
		UserServiceRpcImpl = &UserServiceRpc{}
	})
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

	token, err := auth.GenerateToken(user.Username, user.Id)
	if err != nil {
		out.Reason = FailToCreateToken
		return
	}

	out.JwtToken = token
	out.Reason = Success
	return
}

func (u *UserServiceRpc) UserRegister(ctx context.Context, in *user.UserRegisterRequest, out *user.UserRegisterResponse) (err error) {
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
			Role:     model.User,
			Features: defatultFeaturesJson,
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
