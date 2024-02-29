package db

import (
	"context"
	"fmt"
	"reflect"

	"github.com/cza14h/nino-work/config"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var instance *gorm.DB

type BaseDao[Model any] struct {
	*gorm.DB
}

func (dao *BaseDao[Model]) Create(record *Model) (err error) {
	err = dao.DB.Create(record).Error
	return
}

func (dao *BaseDao[Model]) UpdateById(record Model) (err error) {
	originalStrcut := reflect.TypeOf(record)
	model := reflect.New(originalStrcut).Elem()
	model.FieldByName("Id").Set(reflect.ValueOf(record).FieldByName("Id"))

	err = dao.DB.Model(&model).Updates(record).Error
	return
}

func (dao *BaseDao[Model]) FindByKey(key string, value any) (result *Model, err error) {
	err = dao.Where(fmt.Sprintf("%s = ?", key), value).First(result).Error
	return
}

func ConnectDB(names ...string) *gorm.DB {
	fallbackName := ""
	if len(names) >= 1 {
		fallbackName = names[0]
	} else {
		conf := config.GetConfig()
		fallbackName = conf.System.DbName
	}

	db, err := gorm.Open(sqlite.Open(fallbackName+".db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Fail to connect database")
	}
	instance = db
	return db
}

func NewDBSession(ctx context.Context) *gorm.DB {
	return instance.WithContext(ctx)
}

func InitBaseDao[Model any](ctx context.Context) BaseDao[Model] {
	return BaseDao[Model]{
		DB: NewDBSession(ctx),
	}
}

func Paginate(page, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageNumber, pageSize := page, size
		if pageNumber == 0 {
			pageNumber = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
