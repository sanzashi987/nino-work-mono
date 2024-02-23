package service

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/dao"
	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/apps/canvas-pro/enums"
	"github.com/cza14h/nino-work/pkg/db"
)

type ProjectService struct{}

var projectService *ProjectService

func init() {
	projectService = &ProjectService{}
}

func GetProjectService() *ProjectService {
	return projectService
}

func (p *ProjectService) Create(ctx context.Context, name, groupCode, jsonConfig, useTemplate string) (string, error) {
	projectDao := dao.NewProjectDao(ctx)

	newProject := &model.ProjectModel{
		BaseModel: model.BaseModel{Name: name},
		Version:   enums.DefaultVersion,
		Config:    jsonConfig,
		Code:      enums.CreateCode(enums.PROJECT),
	}
	return newProject.Code, projectDao.Create(newProject)
}

func (p *ProjectService) Update(ctx context.Context) {

}

type ProjectDetailResponse struct {
	Id   string
	Code string
	db.BaseModify
}

func (p *ProjectService) GetInfoById(ctx context.Context, id string) (*ProjectDetailResponse, error) {
	projectDao := dao.NewProjectDao(ctx)

	project, err := projectDao.FindByKey("id", id)
	if err != nil {
		return nil, err
	}

}

type ProjectInfo struct {
	Id        string
	Name      string
	Code      string
	Thumbnail string
	db.BaseModify
}
type ProjectListResponse = []ProjectInfo

func (p *ProjectService) GetList(ctx context.Context, page, size int, name, group, workspace string) (*ProjectListResponse, error) {

	projectDao := dao.NewProjectDao(ctx)

}
