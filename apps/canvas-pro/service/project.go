package service

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/dao"
)

type ProjectService struct{}

var projectService *ProjectService

func init() {
	projectService = &ProjectService{}

}

func GetProjectService() *ProjectService {
	return projectService
}

func (p *ProjectService) Create(ctx context.Context, name, groupCode, jsonConfig, useTemplate string) {
	projectDao := dao.NewProjectDao(ctx)
	projectDao.Create()
}
