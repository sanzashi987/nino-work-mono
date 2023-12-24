package db

import (
	"strconv"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;not null"`
	gorm.Model
}

func (model *BaseModel) GetStringID() string {
	return strconv.FormatUint(model.ID, 10)
}

func (model *BaseModel) GetCreatedAtDate() string {
	return model.CreatedAt.Format("2006-01-02")
}

func (model *BaseModel) GetUpdatedDate() string {
	return model.UpdatedAt.Format("2006-01-02")
}