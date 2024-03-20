package service

import (
	"context"
	"errors"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/dao"
	"github.com/cza14h/nino-work/apps/canvas-pro/enums"
)

type GroupService struct{}

var groupService *GroupService

func init() {
	groupService = &GroupService{}
}

func GetGroupService() *GroupService {
	return groupService
}

var ErrorNameContainIllegalChar = errors.New("error name contains illegal character")

func (serv *GroupService) Create(ctx context.Context, name, workspace string) (err error) {

	if enums.LegalNameReg.FindStringIndex(name) == nil {
		err = ErrorNameContainIllegalChar
		return
	}

	groupDao := dao.NewGroupDao(ctx)

	record, err := groupDao.FindByKey("name", name)
	if record != nil {
	}

	return
}
