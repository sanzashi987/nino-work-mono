package dao

import (
	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/pkg/db"
)

func ConnectDB(dbname string) {
	instance := db.ConnectDB(dbname)
	instance.AutoMigrate(
		&model.ThemeModel{},
		&model.WorkspaceModel{},
		&model.AssetModel{},
		&model.AssetGroupModel{},
		&model.CanvasUserModel{},
		&model.ProjectModel{},
		&model.ProjectGroupModel{},
		&model.TemplateModel{},
		&model.TemplateGroupModel{},
	)
}
