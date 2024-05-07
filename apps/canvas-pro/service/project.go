package service

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/consts"
	"github.com/cza14h/nino-work/apps/canvas-pro/db/dao"
	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/apps/canvas-pro/http/request"
)

type ProjectService struct{}

var ProjectServiceImpl *ProjectService = &ProjectService{}

func (serv ProjectService) Create(ctx context.Context, name, jsonConfig string, groupCode, useTemplate *string) (string, error) {
	projectDao := dao.NewProjectDao(ctx)

	newProject := &model.ProjectModel{}
	newProject.TypeTag, newProject.Name, newProject.Version, newProject.Config = consts.PROJECT, name, consts.DefaultVersion, jsonConfig

	if useTemplate != nil {
		templateDao := dao.NewTemplateDao(ctx)
		templateId, _, _ := consts.GetIdFromCode(*useTemplate)

		template, e := templateDao.FindByKey("id", templateId)
		if e != nil {
			return "", e
		}
		newProject.Config = template.Config
	}

	if groupCode != nil {
		groupId, _, err := consts.GetIdFromCode(*groupCode)
		if err != nil {
			return "", err
		}
		newProject.GroupId = groupId

	}

	if err := projectDao.Create(newProject); err != nil {
		return "", err
	}

	return newProject.Code, nil
}

func (serv ProjectService) Update(ctx context.Context, code string, name, config, thumbnail, group *string) (err error) {
	projectDao := dao.NewProjectDao(ctx)

	toUpdate, idProject := model.ProjectModel{}, model.ProjectModel{}
	var id uint64
	if id, _, err = consts.GetIdFromCode(code); err != nil {
		return
	}

	idProject.Id, toUpdate.Id = id, id
	projectDao.GetOrm().First(&idProject)

	if name != nil {
		toUpdate.Name = *name
	}
	if thumbnail != nil {
		toUpdate.Thumbnail = *thumbnail
	}

	if config != nil {
		toUpdate.Config = *config
	}

	if group != nil {
		var groupId uint64
		if groupId, _, err = consts.GetIdFromCode(*group); err != nil {
			return
		}
		toUpdate.GroupId = groupId
	}

	return projectDao.UpdateById(toUpdate)
}

type ProjectDetail struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	request.DBTime
}

func (serv ProjectService) GetInfoById(ctx context.Context, code string) (result *ProjectDetail, err error) {
	projectDao := dao.NewProjectDao(ctx)
	project, e := projectDao.FindByKey("code", code)
	if e != nil {
		err = e
		return
	}

	result.Code, result.Name, result.Thumbnail = code, project.Name, project.Thumbnail
	result.CreateTime, result.UpdateTime = project.GetCreatedDate(), project.GetUpdatedDate()
	return
}

func (serv ProjectService) LogicalDeletion(ctx context.Context, codes []string) (err error) {
	projectDao := dao.NewProjectDao(ctx)

	intIds := []uint64{}

	for _, code := range codes {
		id, _, _ := consts.GetIdFromCode(code)
		intIds = append(intIds, id)
	}

	return projectDao.BatchLogicalDelete(intIds)
}

type ProjectInfo struct {
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	// Group     string
	request.DBTime
}

// type ProjectInfoList = []ProjectInfo

func (serv ProjectService) GetList(ctx context.Context, workspaceId uint64, page, size int, name, group *string) ([]ProjectInfo, error) {
	projectDao := dao.NewProjectDao(ctx)

	var groupId *uint64

	if group != nil {
		id, _, err := consts.GetIdFromCode(*group)
		if err != nil {
			return nil, err
		}
		groupId = &id
	}

	infos, err := projectDao.GetList(page, size, workspaceId, name, groupId)
	if err != nil {
		return nil, err
	}

	result := []ProjectInfo{}

	for _, info := range *infos {
		temp := ProjectInfo{}
		temp.Name, temp.CreateTime, temp.UpdateTime = info.Name, info.GetCreatedDate(), info.GetUpdatedDate()
		result = append(result, temp)
	}

	return result, nil

}

func (serv ProjectService) PublishProject(ctx context.Context, code string, pulishFlag int, publishSecretKey *string) error {
	return nil
}

func (serv ProjectService) Duplicate(ctx context.Context, code string) (string, error) {
	projectDao := dao.NewProjectDao(ctx)
	copyFromId, _, err := consts.GetIdFromCode(code)
	if err != nil {
		return "", err
	}
	copyFrom, err := projectDao.FindByKey("id", copyFromId)
	name := copyFrom.Name + "_copied"
	hasGroupCode := copyFrom.GroupId == 0
	var groupCode *string = nil
	if hasGroupCode {
		str := consts.GetCodeFromId(consts.GROUP, copyFrom.GroupId)
		groupCode = &str
	}

	return serv.Create(ctx, name, copyFrom.Config, groupCode, nil)
}

func commonMoveGroup(codes []string, groupCode string) (groupId uint64, ids []uint64, err error) {
	groupId, _, err = consts.GetIdFromCode(groupCode)
	if err != nil {
		return
	}

	for _, code := range codes {
		id, _, errInside := consts.GetIdFromCode(code)
		if errInside != nil {
			err = errInside
			return
		}
		ids = append(ids, id)
	}
	return
}

func (serv *ProjectService) BatchMoveGroup(ctx context.Context, workspaceId uint64, projectCodes []string, groupName, groupCode string) error {
	code := groupCode
	projectDao := dao.NewProjectDao(ctx)
	projectDao.BeginTransaction()

	if newGroup, err := createGroup(ctx, (*dao.AnyDao[model.ProjectModel])(projectDao), workspaceId, groupName, consts.PROJECT); err != nil {
		return err
	} else if newGroup != nil {
		code = newGroup.Code
	}

	groupId, projectIds, err := commonMoveGroup(projectCodes, code)
	if err != nil {
		return err
	}

	if err := projectDao.BatchMoveGroup(groupId, workspaceId, projectIds); err != nil {
		return err
	}

	projectDao.CommitTransaction()
	return nil

}

func (serv ProjectService) GetCountFromGroupId(ctx context.Context, workspaceId uint64, groupId []uint64) ([]dao.GroupCount, error) {
	projectDao := dao.NewProjectDao(ctx)

	return projectDao.GetProjectCountByGroup(workspaceId, groupId)
}
