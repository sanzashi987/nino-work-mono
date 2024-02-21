package service

import (
	"context"
	"math"

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
		Name:    name,
		Version: enums.DefaultVersion,
		Config:  jsonConfig,
		Code:    enums.CreateCode(enums.PROJECT),
	}
	return newProject.Code, projectDao.Create(newProject)
}

func (p *ProjectService) Update(ctx context.Context) {

}

type ProjectInfoResponse struct {
	Id   string
	Code string
	db.BaseModify
}

func (p *ProjectService) GetInfoById(ctx context.Context, id string) (*ProjectInfoResponse, error) {
	projectDao := dao.NewProjectDao(ctx)

	project, err := projectDao.FindByKey("id", id)
	if err != nil {
		return nil, err
	}

}

func (p *ProjectService) GetList(ctx context.Context, page, size int, name, group string) (*ProjectInfoResponse, error) {
	pageFallback := int(math.Max(float64(page), 1))
	sizeFallback := int(math.Max(float64(size), 10))
}
