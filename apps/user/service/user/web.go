package userService

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/user/db/dao"
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"github.com/sanzashi987/nino-work/pkg/shared"
)

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

func GetUserInfo(ctx context.Context, userId uint64) (*UserInfoResponse, error) {

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

type UserBio struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
}

type ListUserResponse struct {
	Data []*UserBio `json:"data"`
	shared.PaginationResponse
}

func ListUser(ctx context.Context, pagination *shared.PaginationRequest) (*ListUserResponse, error) {
	tx := db.NewTx(ctx)

	var users []*model.UserModel

	var count int64

	if err := tx.Model(&model.UserModel{}).Count(&count).Error; err != nil {
		return nil, err
	}

	if err := tx.Scopes(db.Paginate(pagination.Page, pagination.Size)).Order("id DESC").Find(&users).Error; err != nil {
		return nil, err
	}

	res := []*UserBio{}
	for _, user := range users {
		res = append(res, &UserBio{
			Id:       user.Id,
			Username: user.Username,
		})
	}

	return &ListUserResponse{
		Data: res,
		PaginationResponse: shared.PaginationResponse{
			PageIndex:   pagination.Page,
			PageSize:    pagination.Size,
			RecordTotal: int(count),
		},
	}, nil

}

type BindRoleRequest struct {
	UserId  uint64   `json:"user_id"`
	RoleIds []uint64 `json:"role_ids"`
}

func BindUserRoles(ctx context.Context, operator uint64, payload *BindRoleRequest) error {
	result, err := GetUserAdmins(ctx, operator)
	if err != nil {
		return err
	}

	if !result.HasAnyAdmin() {
		return ErrNopermission
	}

	hasPermission, err := result.ToPermissionSet()
	if err != nil {
		return err
	}

	tx := result.Tx

	tryToUsePermission, err := dao.FindAllPermissionsWithRoleIds(tx, payload.RoleIds)

	if err != nil {
		return err
	}

	if !hasPermission.IsStrictlyContains(tryToUsePermission) {
		return ErrOutsidepermission
	}

	toBindUser := &model.UserModel{}
	toBindUser.Id = payload.UserId
	toBindRoles := []*model.RoleModel{}
	for _, roleId := range payload.RoleIds {
		role := &model.RoleModel{}
		role.Id = roleId
		toBindRoles = append(toBindRoles, role)
	}

	if err := tx.Model(toBindUser).Association("Roles").Replace(toBindRoles); err != nil {
		return err
	}

	return nil
}
