package db

import (
	"strconv"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint64 `gorm:"column:id;primaryKey;not null;index:,unique;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
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

// func (model *BaseModel) CreateBaseModel() *BaseModel {
// 	return &BaseModel{

// 	}
// }
