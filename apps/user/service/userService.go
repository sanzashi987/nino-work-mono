package service

import (
	"context"
	"errors"
	"sync"

	"github.com/cza14h/nino-work/apps/user/db/dao"
	"github.com/cza14h/nino-work/apps/user/db/model"
	"github.com/cza14h/nino-work/pkg/auth"
	"github.com/cza14h/nino-work/proto/user"
	"gorm.io/gorm"
)

type UserServiceRpcImpl struct{}

var UserServiceRpc *UserServiceRpcImpl
var once sync.Once

func GetUserServiceRpc() *UserServiceRpcImpl {
	once.Do(func() {
		UserServiceRpc = &UserServiceRpcImpl{}
	})
	return UserServiceRpc
}

func (u *UserServiceRpcImpl) UserLogin(ctx context.Context, in *user.UserLoginRequest, out *user.UserLoginResponse) (err error) {
	user, err := dao.NewUserDao(ctx).FindUserByUsername(in.Username)
	if err != nil {
		out.Reason = UsernameNotExist
		return
	}

	if valid := user.CheckPassowrd(in.Password); !valid {
		out.Reason = PasswordNotMatch
		return
	}

	token, err := auth.GenerateToken(user.Username, user.ID)
	if err != nil {
		out.Reason = FailToCreateToken
		return
	}

	out.JwtToken = token
	out.Reason = Success
	return
}

func (u *UserServiceRpcImpl) UserRegister(ctx context.Context, in *user.UserRegisterRequest, out *user.UserRegisterResponse) error {
	dbSession := dao.NewUserDao(ctx)
	user, err := dbSession.FindUserByUsername(in.Username)
	if user != nil {
		out.Reason = UsernameExisted
		return err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {

		if in.Password != in.PasswordConfirm {
			out.Reason = PasswordNotMatch
			return err
		}

		user := model.UserModel{
			Username: in.Username,
			Password: in.Password,
			Fobidden: false,
			Role:     model.User,
			Features: defatultFeaturesJson,
		}

		dbSession.CreateUser(&user)
		token, err := auth.GenerateToken(user.Username, user.ID)
		if err != nil {
			out.Reason = FailToCreateToken
			return err
		}
		out.JwtToken = token
		return nil
	}

	out.Reason = InternalServerError
	return errors.New("Unknown edge case in user service")
}
