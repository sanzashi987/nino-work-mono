package db

import (
	"strconv"
	"time"

	"gorm.io/gorm"
)

type BaseTime struct {
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

type BaseModel struct {
	BaseTime
	Id         uint64         `gorm:"column:id;primaryKey;not null;index;unique;"`
	DeleteTime gorm.DeletedAt `gorm:"index"`
	IGetId     `gorm:"-"`
}

type IGetId interface {
	GetId() uint64
}

type GetDeleteTime interface {
	GetDeleteTime() *time.Time
}

func (model *BaseModel) GetId() uint64 {
	return model.Id
}

func (model BaseModel) GetDeleteTime() *time.Time {
	return &model.DeleteTime.Time
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

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreateTime = time.Now()
	m.UpdateTime = m.CreateTime
	return
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdateTime = time.Now()
	return
}

func ToIdList(a []IGetId) []uint64 {
	res := make([]uint64, len(a))
	for i, model := range a {
		res[i] = model.GetId()
	}
	return res
}
