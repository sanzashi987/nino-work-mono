package project

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/dao"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

func GetInfoById(ctx context.Context, workspaceId uint64, code string) (*ProjectDetail, error) {

	result, project := ProjectDetail{}, model.ProjectModel{}
	tx := db.NewTx(ctx)

	if err := tx.Where("code = ? ", code).Find(&project).Error; err != nil {
		return nil, err
	}

	result.Code, result.Name, result.Thumbnail = code, project.Name, project.Thumbnail
	result.CreateTime, result.UpdateTime = project.GetCreatedDate(), project.GetUpdatedDate()
	return &result, nil
}

func List(ctx context.Context, workspaceId uint64, page, size int, name, group *string) ([]*ProjectInfo, error) {
	tx := db.NewTx(ctx)

	var groupId *uint64

	if group != nil {
		id, _, err := consts.GetIdFromCode(*group)
		if err != nil {
			return nil, err
		}
		groupId = &id
	}

	infos, err := dao.GetList(tx, page, size, workspaceId, name, groupId)
	if err != nil {
		return nil, err
	}

	result := []*ProjectInfo{}

	for _, info := range *infos {
		temp := &ProjectInfo{}
		temp.Name, temp.CreateTime, temp.UpdateTime = info.Name, info.CreateTime.Unix(), info.CreateTime.Unix()
		temp.Code = info.Code
		result = append(result, temp)
	}

	return result, nil

}
