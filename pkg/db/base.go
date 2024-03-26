package db

import (
	"strconv"
	"time"
)

type BaseTime struct {
	CreateTime time.Time
	UpdateTime time.Time
}

type BaseModel struct {
	BaseTime
	Id         uint64     `gorm:"column:id;primaryKey;not null;index;unique;"`
	DeleteTime *time.Time `gorm:"index"`
}

func (model *BaseModel) GetStringID() string {
	return strconv.FormatUint(model.Id, 10)
}

func (model *BaseModel) GetCreatedDate() string {
	return model.CreateTime.Format("2006-01-02")
}

func (model *BaseModel) GetUpdatedDate() string {
	return model.UpdateTime.Format("2006-01-02")
}
