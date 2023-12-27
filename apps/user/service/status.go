package service

const (
	Success             int32 = 0
	UsernameNotExist          = 1
	UsernameExisted           = 2
	PasswordNotMatch          = 3
	FailToCreateToken         = 4
	InternalServerError       = 505
)
