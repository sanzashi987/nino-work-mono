package service

const (
	RegisterSuccess     int32 = 0
	UsernameExisted           = 1
	PasswordNotMatch          = 2
	FailToCreateToken         = 3
	InternalServerError       = 505
)
