package group

import (
	"context"
	"sort"

	"github.com/sanzashi987/nino-work/apps/canvix/db/dao"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"github.com/sanzashi987/nino-work/pkg/shared"
)

type ListGroupOutput struct {
	Id    uint64 `json:"id"`
	Name  string `json:"name"`
	Code  string `json:"code"`
	Count uint   `json:"count"`
}

type ListGroupOutputs []ListGroupOutput

func (p ListGroupOutputs) Len() int {
	return len(p)
}
func (p ListGroupOutputs) Less(i, j int) bool {
	return p[i].Id < p[j].Id
}
func (p ListGroupOutputs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type GroupCount struct {
	Id    uint64 `gorm:"column:id"`
	Count uint64 `gorm:"column:count"`
}


type ListGroupReq struct {
	GroupCode string `json:"code"`
	GroupName string `json:"name"`
	TypeTag   string `json:"type" binding:"required"`
	shared.PaginationRequest
}

func List(ctx context.Context, workspaceId uint64, req *ListGroupReq) (ListGroupOutputs, error) {

	tx := db.NewTx(ctx)

	records, err := dao.FindByNameAndWorkspace(tx, workspaceId, req.GroupName, req.TypeTag)
	if err != nil {
		return nil, err
	}

	m, exist := typeTagToGroupCountHandler[req.TypeTag]
	if !exist {
		err = errTagNotSupported
		return nil, err
	}

	groupIds := []uint64{}

	idToRecord := map[uint64]*model.GroupModel{}
	for _, record := range records {
		groupIds = append(groupIds, record.Id)
		idToRecord[record.Id] = record
	}

	groupCounts := []*GroupCount{}

	if err := tx.Model(&m).Where("workspace = ? AND group_id IN ?", workspaceId, groupIds).Select("id", "COUNT(id) as count").Group("group_id").Find(&groupCounts).Error; err != nil {
		return nil, err
	}

	output := ListGroupOutputs{}
	for _, groupCount := range groupCounts {
		record := idToRecord[groupCount.Id]
		output = append(output, ListGroupOutput{
			Id:    groupCount.Id,
			Count: uint(groupCount.Count),
			Name:  record.Name,
			Code:  record.Code,
		})
	}

	sort.Sort(output)

	return output, nil
}
