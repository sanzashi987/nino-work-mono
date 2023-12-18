package service

import (
	"context"
	"sync"

	"github.com/cza14h/nino-work/proto/user"
)

type UserServiceRpcImpl struct {
}

var UserServiceRpc *UserServiceRpcImpl
var once *sync.Once

func GetUserServiceRpc() *UserServiceRpcImpl {
	once.Do(func() {
		UserServiceRpc = &UserServiceRpcImpl{}
	})
	return UserServiceRpc
}

func (u *UserServiceRpcImpl) UserLogin(ctx context.Context, in *user.UserLoginRequest, out *user.UserLoginResponse) (err error) {

	return
}

func (u *UserServiceRpcImpl) UserRegister(ctx context.Context, in *user.UserRegisterRequest, out *user.UserLoginResponse) (err error) {

	return
}
