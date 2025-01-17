package service

import (
	"context"
	"errors"
	"time"

	"github.com/sanzashi987/nino-work/apps/user/db/dao"
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/controller"
	"github.com/sanzashi987/nino-work/pkg/db"
	"github.com/sanzashi987/nino-work/proto/user"
	"gorm.io/gorm"
)

type UserServiceRpc struct{}

var UserServiceRpcImpl *UserServiceRpc = &UserServiceRpc{}

func GetUserServiceRpc() user.UserServiceHandler {
	return UserServiceRpcImpl
}

func (u *UserServiceRpc) UserLogin(ctx context.Context, in *user.UserLoginRequest, out *user.UserLoginResponse) (err error) {
	user, err := dao.FindUserByUsername(db.NewTx(ctx), in.Username)
	if err != nil {
		out.Reason = UsernameNotExist
		return
	}

	if valid := user.CheckPassowrd(in.Password); !valid {
		err = errors.New("user password not match")
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

	token, err = controller.GenerateToken(user.Username, user.Id, time.Hour*24*time.Duration(days))
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
	tx := db.NewTx(ctx)

	user, err := dao.FindUserByUsername(tx, in.Username)
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

		dao.CreateUser(tx, &user)
		var token string
		token, err = controller.GenerateToken(user.Username, user.Id)
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
	Name      string `json:"name"`
	Code      string `json:"code"`
	Icon      string `json:"icon"`
	Hyperlink bool   `json:"hyperlink"`
	Path      string `json:"path"`
	Type      uint8  `json:"type"`
	Order     int    `json:"order"`
}

type UserInfoResponse struct {
	UserId      uint64      `json:"user_id"`
	Username    string      `json:"username"`
	Menus       []*MenuMeta `json:"menus"`
	Permissions []*CodeName `json:"permissions"`
	Roles       []*CodeName `json:"roles"`
}

func (u *UserServiceWeb) GetUserInfo(ctx context.Context, userId uint64) (*UserInfoResponse, error) {

	user, tx, err := getUserRolePermission(ctx, userId)
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
				Name:      menu.Name,
				Code:      code,
				Icon:      menu.Icon,
				Hyperlink: menu.Hyperlink,
				Path:      menu.Path,
				Order:     menu.Order,
				Type:      uint8(menu.Type),
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

func (u *UserServiceWeb) GetUserRoles(ctx context.Context, userId uint64) ([]model.RoleModel, error) {

	user, err := dao.FindUserWithRoles(db.NewTx(ctx), userId)
	if err != nil {
		return nil, err
	}

	return user.Roles, nil
}

func (u *UserServiceWeb) GetUserRoleWithPermissions(ctx context.Context, userId uint64) (*model.UserModel, error) {
	user, _, err := getUserRolePermission(ctx, userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func getUserRolePermission(ctx context.Context, userId uint64) (*model.UserModel, *gorm.DB, error) {
	tx := db.NewTx(ctx)
	user, err := dao.FindUserWithRoles(tx, userId)
	if err != nil {
		return nil, nil, err
	}

	userRoles := []*model.RoleModel{}
	for _, role := range user.Roles {
		userRoles = append(userRoles, &role)
	}

	err = dao.FindRolesWithPermissions(tx, userRoles...)

	if err != nil {
		return nil, nil, err
	}

	res := []model.RoleModel{}

	for _, role := range userRoles {
		res = append(res, *role)
	}

	user.Roles = res

	return user, tx, nil
}

type UserAdminResult struct {
	SuperAdminApps []*model.ApplicationModel
	AdminApps      []*model.ApplicationModel
}

func getUserAdmins(ctx context.Context, userId uint64) (*UserAdminResult, error) {
	user, tx, err := getUserRolePermission(ctx, userId)
	if err != nil {
		return nil, err
	}

	applications := map[uint64]bool{}
	permissions := map[uint64]bool{}
	for _, role := range user.Roles {
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
	err = tx.Table("applications").Where("id IN ?", appIds).Find(&apps).Error
	if err != nil {
		return nil, err

	}

	superRes := []*model.ApplicationModel{}
	adminRes := []*model.ApplicationModel{}
	superResMap := map[uint64]*model.ApplicationModel{}
	adminResMap := map[uint64]*model.ApplicationModel{}
	for _, app := range apps {
		if _, superExist := permissions[app.SuperAdmin]; superExist {
			superResMap[app.Id] = &app
		}
	}

	for _, app := range apps {
		if _, adminExist := permissions[app.Admin]; adminExist {
			adminResMap[app.Id] = &app
		}
	}

	for _, app := range superResMap {
		superRes = append(superRes, app)
	}
	for _, app := range adminResMap {
		adminRes = append(adminRes, app)
	}
	result := UserAdminResult{
		SuperAdminApps: superRes,
		AdminApps:      adminRes,
	}

	return &result, nil
}
