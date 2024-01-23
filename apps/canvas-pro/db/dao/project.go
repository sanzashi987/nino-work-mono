package dao

import (
	"context"

	"gorm.io/gorm"
)

type ProjectDao struct {
	*gorm.DB
}

func NewProjectDao(ctx context.Context) *ProjectDao {
	return &ProjectDao{
		DB: newDBSession(ctx),
	}
}


func (p *ProjectDao) CreateProject(){
	
}