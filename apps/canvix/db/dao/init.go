package dao

import (
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type AnyDao[T any] struct {
	db.BaseDao[T]
}

func ConnectDB() {
	instance := db.ConnectDB()
	instance.AutoMigrate(
		&model.ThemeModel{},
		&model.WorkspaceModel{},
		&model.AssetModel{},
		&model.CanvasUserModel{},
		&model.ProjectModel{},
		&model.GroupModel{},
		&model.TemplateModel{},
	)
}
