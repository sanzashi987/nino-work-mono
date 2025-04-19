package userService

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/user/db/dao"
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"github.com/sanzashi987/nino-work/pkg/shared"
)

type MenuMeta struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Icon  string `json:"icon"`
	Path  string `json:"path"`
	Type  uint8  `json:"type"`
	Order int    `json:"order"`
}

type UserInfoResponse struct {
	UserId      uint64             `json:"user_id"`
	Username    string             `json:"username"`
	Menus       []*MenuMeta        `json:"menus"`
	Permissions []*shared.EnumMeta `json:"permissions"`
	Roles       []*shared.EnumMeta `json:"roles"`
}

func GetUserInfo(ctx context.Context, userId uint64) (*UserInfoResponse, error) {

	user, tx, err := GetUserRolePermission(ctx, userId)
	if err != nil {
		return nil, err
	}
	resRoles := []*shared.EnumMeta{}
	resPermissions := []*shared.EnumMeta{}
	permissions := map[uint64]bool{}
	for _, role := range user.Roles {
		resRoles = append(resRoles, &shared.EnumMeta{
			Name:  role.Name,
			Value: role.Code,
		})
		for _, permission := range role.Permissions {
			permissions[permission.Id] = true
			resPermissions = append(resPermissions, &shared.EnumMeta{
				Name:  permission.Name,
				Value: permission.Code,
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

type ListUserResponse = shared.ResponseWithPagination[[]*UserBio]

// TODO specific permission check for the operator
func ListUser(ctx context.Context, pagination *shared.PaginationRequest) (*ListUserResponse, error) {
	tx := db.NewTx(ctx)

	r, err := db.QueryWithTotal[model.UserModel](tx.Model(&model.UserModel{}), pagination.Page, pagination.Size)
	if err != nil {
		return nil, err
	}

	data := []*UserBio{}
	for _, user := range r.Records {
		data = append(data, &UserBio{
			Id:       user.Id,
			Username: user.Username,
		})
	}

	res := &ListUserResponse{}
	res.Init(data, r.Page, r.Total)

	return res, nil

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

// TODO specific permission check for the operator
func GetUserRoles(ctx context.Context, user, targetUser uint64) ([]*shared.EnumMeta, error) {

	targetUserModel := &model.UserModel{}
	targetUserModel.Id = targetUser

	tx := db.NewTx(ctx)
	if err := tx.Model(targetUserModel).Association("Roles").Find(&targetUserModel.Roles); err != nil {
		return nil, err
	}

	result := []*shared.EnumMeta{}
	for _, role := range targetUserModel.Roles {
		result = append(result, &shared.EnumMeta{
			Name:  role.Name,
			Value: role.Id,
		})
	}
	return result, nil
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUserByAdmin(ctx context.Context, user uint64, payload *CreateUserRequest) (uint64, error) {
	result, err := GetUserAdmins(ctx, user)
	if err != nil {
		return 0, err
	}

	if !result.HasAnyAdmin() {
		return 0, ErrNopermission
	}

	tx := result.Tx

	userModel := &model.UserModel{
		Username: payload.Username,
		Password: payload.Password,
	}

	if err := dao.CreateUser(tx, userModel); err != nil {
		return 0, err
	}

	return userModel.Id, nil
}
