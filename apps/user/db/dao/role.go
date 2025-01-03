package dao

import (
	"context"
	"errors"

	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type RoleDao struct {
	db.BaseDao[model.RoleModel]
}

func NewRoleDao(ctx context.Context, dao ...*db.BaseDao[model.RoleModel]) *RoleDao {
	return &RoleDao{BaseDao: db.NewDao(ctx, dao...)}
}

func (dao *RoleDao) FindRolesWithPermissions(roles ...*model.RoleModel) error {
	if len(roles) == 0 {
		return errors.New("roles is required")
	}

	// 收集所有角色ID
	roleIds := []uint64{}
	roleMap := map[uint64]*model.RoleModel{}
	for i := range roles {
		roleIds = append(roleIds, roles[i].Id)
		roleMap[roles[i].Id] = roles[i]
	}

	// 一次性查询所有角色及其权限
	rolesWithPerms := []*model.RoleModel{}
	err := dao.GetOrm().
		Preload("Permissions").
		Where("id IN ?", roleIds).
		Find(&rolesWithPerms).Error
	if err != nil {
		return err
	}

	// 将查询结果写回原数组
	for _, role := range rolesWithPerms {
		if r, exists := roleMap[role.Id]; exists {
			r.Permissions = role.Permissions
		}
	}

	return nil
}
