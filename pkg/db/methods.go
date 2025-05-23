package db

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"gorm.io/gorm"
)

type ORMConfig struct {
	TableName string
}

type Configure = func(*ORMConfig)

func TableName(tableName string) Configure {
	return func(o *ORMConfig) {
		o.TableName = tableName
	}
}

func (dao *BaseDao[Model]) GetOrm(config ...Configure) *gorm.DB {
	defaultConfig := ORMConfig{}

	for _, fc := range config {
		fc(&defaultConfig)
	}

	orm := dao.db
	if dao.transaction != nil {
		orm = dao.transaction
	}

	if defaultConfig.TableName != "" {
		orm = orm.Table(defaultConfig.TableName)
	}

	return orm
}

func (dao *BaseDao[Model]) CommitTransaction() {
	if dao.transaction != nil {
		dao.transaction.Commit()
		dao.transaction = nil
	}
}

func (dao *BaseDao[Model]) Create(record *Model, config ...Configure) (err error) {
	err = dao.GetOrm(config...).Create(record).Error
	return
}

var ErrorIdIsNotProvided = errors.New("id is not provided")
var ErrorRecordIsNotAStructInstance = errors.New("record is not a struct instance or instance pointer")

func (dao *BaseDao[Model]) UpdateById(record Model, config ...Configure) (err error) {
	originalStruct := reflect.TypeOf(record)
	var structType reflect.Type = originalStruct

	if originalStruct.Kind() == reflect.Ptr {
		if originalStruct.Elem().Kind() == reflect.Struct {
			structType = originalStruct.Elem()
		} else {
			err = ErrorRecordIsNotAStructInstance
			return
		}
	}

	model := reflect.New(structType).Elem()
	if model.FieldByName("Id").IsZero() {
		err = ErrorIdIsNotProvided
		return
	}
	model.FieldByName("Id").Set(reflect.ValueOf(record).FieldByName("Id"))

	err = dao.GetOrm(config...).Model(model.Elem().Interface()).Updates(record).Error
	return
}

var ErrorNotADbModel = errors.New("the record is not a db model")

func (dao *BaseDao[Model]) LogicalDelete(record Model, config ...Configure) (err error) {
	originalStruct := reflect.TypeOf(record)
	reflectRecord := reflect.ValueOf(record)
	reflectId := reflectRecord.FieldByName("Id")
	if reflectId.IsZero() {
		err = ErrorIdIsNotProvided
		return
	}

	model := reflect.New(originalStruct)
	model.FieldByName("Id").Set(reflectId)

	return dao.GetOrm(config...).Model(model.Elem().Interface()).Update("delete_time", time.Now()).Error

}

func (dao *BaseDao[Model]) FindByKey(key string, value any, config ...Configure) (result *Model, err error) {
	err = dao.GetOrm(config...).Where(fmt.Sprintf("%s = ?", key), value).First(result).Error
	return
}

func CommonSuggest[T any](ctx context.Context, keyword string, results *[]T) error {
	tx := NewTx(ctx)
	if err := tx.
		Where("name LIKE ? OR code LIKE ?", "%"+keyword+"%", "%"+keyword+"%").
		Find(results).Error; err != nil {
		return err
	}

	return nil
}
