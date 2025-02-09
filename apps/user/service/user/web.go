package userService

import (
	"context"

	"github.com/sanzashi987/nino-work/proto/user"
)

// GetApplicationPermissions implements user.UserServiceHandler.
func (u *UserServiceRpc) GetApplicationPermissions(context.Context, *user.ApplicationPermissionsRequest, *user.ApplicationPermissionsResponse) error {
	panic("unimplemented")
}

// GetUserPermissions implements user.UserServiceHandler.
func (u *UserServiceRpc) GetUserPermissions(context.Context, *user.UserPermissionsRequest, *user.UserPermissionsResponse) error {
	panic("unimplemented")
}

type UserServiceWeb struct{}

var UserServiceWebImpl *UserServiceWeb = &UserServiceWeb{}

type CodeName struct {
	Name string `json:"name"`
	Code string `json:"code"`
}
type MenuMeta struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Icon  string `json:"icon"`
	Path  string `json:"path"`
	Type  uint8  `json:"type"`
	Order int    `json:"order"`
}

type UserInfoResponse struct {
	UserId      uint64      `json:"user_id"`
	Username    string      `json:"username"`
	Menus       []*MenuMeta `json:"menus"`
	Permissions []*CodeName `json:"permissions"`
	Roles       []*CodeName `json:"roles"`
}

func (u *UserServiceWeb) GetUserInfo(ctx context.Context, userId uint64) (*UserInfoResponse, error) {

	user, tx, err := GetUserRolePermission(ctx, userId)
	if err != nil {
		return nil, err
	}
	resRoles := []*CodeName{}
	resPermissions := []*CodeName{}
	permissions := map[uint64]bool{}
	for _, role := range user.Roles {
		resRoles = append(resRoles, &CodeName{
			Name: role.Name,
			Code: role.Code,
		})
		for _, permission := range role.Permissions {
			permissions[permission.Id] = true
			resPermissions = append(resPermissions, &CodeName{
				Name: permission.Name,
				Code: permission.Code,
			})
		}
	}

	if err := tx.Preload("Menus").Find(&user.Roles).Error; err != nil {
		return nil, err
	}

	menuMap := map[string]*MenuMeta{}
	for _, role := range user.Roles {
		for _, menu := range role.Menus {
			code := menu.Code
			if _, exist := menuMap[code]; exist {
				continue
			}
			menuMap[code] = &MenuMeta{
				Name:  menu.Name,
				Code:  code,
				Icon:  menu.Icon,
				Path:  menu.Path,
				Order: menu.Order,
				Type:  uint8(menu.Type),
			}

		}
	}

	resMenus := []*MenuMeta{}
	for _, menu := range menuMap {
		resMenus = append(resMenus, menu)
	}

	return &UserInfoResponse{
		UserId:      user.Id,
		Username:    user.Username,
		Permissions: resPermissions,
		Menus:       resMenus,
		Roles:       resRoles,
	}, nil

}

