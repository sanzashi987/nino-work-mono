package db

import (
	"context"
	"errors"
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

func (dao *BaseDao[Model]) Create(record *Model, table ...string) (err error) {
	orm := dao.DB
	if len(table) > 0 {
		orm = orm.Table(table[0])
	}
	err = orm.Create(record).Error
	return
}

var ErrorIdIsNotProvided = errors.New("id is not provided")

func (dao *BaseDao[Model]) UpdateById(record Model, table ...string) (err error) {
	originalStruct := reflect.TypeOf(record)
	model := reflect.New(originalStruct).Elem()
	if model.FieldByName("Id").IsZero() {
		err = ErrorIdIsNotProvided
		return
	}
	model.FieldByName("Id").Set(reflect.ValueOf(record).FieldByName("Id"))

	orm := dao.DB
	if len(table) > 0 {
		orm = orm.Table(table[0])
	}
	err = orm.Model(&model).Updates(record).Error
	return
}

var ErrorNotADbModel = errors.New("the record is not a db model")

func (dao *BaseDao[Model]) LogicalDelete(record Model, table ...string) (err error) {

	model := reflect.ValueOf(record)

	var tableName = model.MethodByName("TableName").Call()
	orm := dao.DB
	if len(table) > 0 {
		orm = orm.Table(table[0])
	}
	return dao.DB.Table(model.ProjectModel{}.TableName()).Where("id IN ?", ids).Update("deleted", model.Deleted).Error

}

func (dao *BaseDao[Model]) FindByKey(key string, value any, table ...string) (result *Model, err error) {
	orm := dao.DB
	if len(table) > 0 {
		orm = orm.Table(table[0])
	}
	err = orm.Where(fmt.Sprintf("%s = ?", key), value).First(result).Error
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
