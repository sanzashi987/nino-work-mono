package service

import (
	"context"
	"errors"

	"github.com/sanzashi987/nino-work/apps/user/db/dao"
	"github.com/sanzashi987/nino-work/apps/user/db/model"
)

type RoleServiceWeb struct{}

var RoleServiceWebImpl *RoleServiceWeb = &RoleServiceWeb{}

// CreateRoleRequest 创建角色请求参数
type CreateRoleRequest struct {
	RoleName        string   `json:"role_name"`
	RoleCode        string   `json:"role_code"`
	RoleDescription string   `json:"role_description"`
	PermissionIds   []uint64 `json:"permission_ids"`
}

// 创建角色
func (r *RoleServiceWeb) CreateRole(ctx context.Context, userId uint64, payload CreateRoleRequest) error {
	userAdmins, roleDao, err := getUserRolePermission(ctx, userId)
	if err != nil {
		return err
	}

	if len(userAdmins) == 0 {
		return errors.New("user does not have any admin permission")
	}

	if payload.RoleName == "" {
		return errors.New("角色编码不能为空")
	}

	if payload.RoleCode == "" {
		return errors.New("角色编码不能为空")
	}

	roleDao.BeginTransaction()

	// 创建角色
	newRole := &model.RoleModel{
		Name:        payload.RoleName,
		Code:        payload.RoleCode,
		Description: payload.RoleDescription,
	}

	if err := roleDao.Create(newRole); err != nil {
		roleDao.RollbackTransaction()
		return err
	}

	// 关联权限
	if len(payload.PermissionIds) > 0 {
		permissions := make([]model.PermissionModel, 0)

		for _, pid := range payload.PermissionIds {
			permission := model.PermissionModel{}
			permission.Id = pid
			permissions = append(permissions, permission)
		}

		if err := roleDao.GetOrm().Model(newRole).Association("Permissions").Replace(permissions); err != nil {
			roleDao.RollbackTransaction()
			return err
		}
	}

	roleDao.CommitTransaction()
	return nil
}

// 获取角色详情
func (r *RoleServiceWeb) GetRoleDetail(ctx context.Context, roleId uint64) (*model.RoleModel, error) {

	roleDao := dao.NewRoleDao(ctx)
	role := model.RoleModel{}
	role.Id = roleId

	if err := roleDao.FindRolesWithPermissions(&role); err != nil {
		return nil, err
	}

	return &role, nil
}

// 更新角色
func (r *RoleServiceWeb) UpdateRole(ctx context.Context, roleId uint64, payload CreateRoleRequest) error {

	roleDao := dao.NewRoleDao(ctx)
	roleDao.BeginTransaction()

	role := &model.RoleModel{
		Name:        payload.RoleName,
		Code:        payload.RoleCode,
		Description: payload.RoleDescription,
	}
	role.Id = roleId

	if err := roleDao.GetOrm().Updates(role).Error; err != nil {
		roleDao.RollbackTransaction()
		return err
	}

	// 更新权限关联
	if len(payload.PermissionIds) > 0 {
		permissions := make([]model.PermissionModel, 0)
		for _, pid := range payload.PermissionIds {
			permission := model.PermissionModel{}
			permission.Id = pid
			permissions = append(permissions, permission)
		}

		if err := roleDao.GetOrm().Model(role).Association("Permissions").Replace(permissions); err != nil {
			roleDao.RollbackTransaction()
			return err
		}
	}

	roleDao.CommitTransaction()
	return nil
}

// 删除角色
func (r *RoleServiceWeb) DeleteRole(ctx context.Context, roleId uint64) error {
	roleDao := dao.NewRoleDao(ctx)
	roleDao.BeginTransaction()

	role := &model.RoleModel{}
	role.Id = roleId

	// 先清除角色关联的权限
	if err := roleDao.GetOrm().Model(role).Association("Permissions").Clear(); err != nil {
		roleDao.RollbackTransaction()
		return err
	}

	// 再删除角色本身
	if err := roleDao.GetOrm().Delete(role).Error; err != nil {
		roleDao.RollbackTransaction()
		return err
	}

	roleDao.CommitTransaction()
	return nil
}

// 模糊搜索角色
func (r *RoleServiceWeb) SuggestRoles(ctx context.Context, keyword string) ([]model.RoleModel, error) {
	roleDao := dao.NewRoleDao(ctx)
	var roles []model.RoleModel

	if err := roleDao.GetOrm().
		Where("name LIKE ? OR code LIKE ?", "%"+keyword+"%", "%"+keyword+"%").
		Find(&roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
}
