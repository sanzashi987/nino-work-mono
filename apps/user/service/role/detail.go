package roleService

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/user/db/dao"
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

func GetRoleDetail(ctx context.Context, roleId uint64) (*model.RoleModel, error) {

	role := model.RoleModel{}
	role.Id = roleId

	if err := dao.FindRolesWithPermissions(db.NewTx(ctx), &role); err != nil {
		return nil, err
	}

	return &role, nil
}
