package db

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func (dao *BaseDao[Model]) Create(record Model, table ...string) (err error) {
	orm := dao.DB
	if len(table) > 0 {
		orm = orm.Table(table[0])
	}
	err = orm.Create(&record).Error
	return
}

var ErrorIdIsNotProvided = errors.New("id is not provided")
var ErrorRecordIsNotAStructInstance = errors.New("record is not a struct instance or instance pointer")

func (dao *BaseDao[Model]) UpdateById(record Model, table ...string) (err error) {
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

	orm := dao.DB
	if len(table) > 0 {
		orm = orm.Table(table[0])
	}
	err = orm.Model(model.Elem().Interface()).Updates(record).Error
	return
}

var ErrorNotADbModel = errors.New("the record is not a db model")

func (dao *BaseDao[Model]) LogicalDelete(record Model, table ...string) (err error) {

	reflectRecord := reflect.ValueOf(record)

	var tableName string

	if len(table) > 0 {
		tableName = table[0]
	} else {
		shouldTableName := reflectRecord.MethodByName("TableName")
		if !shouldTableName.IsValid() {
			modelStruct := reflect.TypeOf(record).Elem().Name()
			tableName = strings.ToLower(modelStruct + "s")
		} else {
			tableName = shouldTableName.Call(nil)[0].String()
		}
	}

	return dao.DB.Table(tableName).Model(record).Update("deleted", Deleted).Error

}

func (dao *BaseDao[Model]) FindByKey(key string, value any, table ...string) (result *Model, err error) {
	orm := dao.DB
	if len(table) > 0 {
		orm = orm.Table(table[0])
	}
	err = orm.Where(fmt.Sprintf("%s = ?", key), value).First(result).Error
	return
}