package project

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"github.com/sanzashi987/nino-work/pkg/shared"
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

type GetProjectListRequest struct {
	shared.PaginationRequest
	// Workspace string
	Name  *string `json:"name"`
	Group *string `json:"group"`
}

type GetProjectListResponse = shared.ResponseWithPagination[[]*ProjectInfo]

func List(ctx context.Context, workspaceId uint64, req *GetProjectListRequest) (*GetProjectListResponse, error) {
	page, size, name, group := req.Page, req.Size, req.Name, req.Group
	tx := db.NewTx(ctx)

	var groupId *uint64

	if group != nil {
		id, _, err := consts.GetIdFromCode(*group)
		if err != nil {
			return nil, err
		}
		groupId = &id
	}

	query := tx.Model(&model.ProjectModel{})
	if groupId != nil {
		query = query.Where(" group_id = ?", *groupId)
	}

	if name != nil {
		query = query.Where(" name LIKE ?", *name)
	}

	r, err := db.QueryWithTotal[model.ProjectModel](query, page, size)
	if err != nil {
		return nil, err
	}

	data := []*ProjectInfo{}

	for _, info := range r.Records {
		temp := &ProjectInfo{}
		temp.Name, temp.CreateTime, temp.UpdateTime = info.Name, info.CreateTime.Unix(), info.CreateTime.Unix()
		temp.Code = info.Code
		data = append(data, temp)
	}

	res := &GetProjectListResponse{}
	res.Init(data, r.Page, r.Total)

	return res, nil

}
