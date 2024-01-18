package model

import "github.com/cza14h/nino-work/pkg/db"

type ProjectGroup struct {
	db.BaseModel
	Workspace string
}

type ProjectModel struct {
	db.BaseModel
}
