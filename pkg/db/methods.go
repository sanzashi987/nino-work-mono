package db

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

type ORMConfig struct {
	TableName string
}

func ConfigureTable(tableName string) func(*ORMConfig) {
	return func(o *ORMConfig) {
		o.TableName = tableName
	}
}

func (dao *BaseDao[Model]) BeginTransaction() {
	dao.transaction = dao.db.Begin()
}

func (dao *BaseDao[Model]) RollbackTransaction() {
	if dao.transaction != nil {
		dao.transaction.Rollback()
	}
}

func (dao *BaseDao[Model]) RollbackToTransaction(name string) {
	if dao.transaction != nil {
		dao.transaction.RollbackTo(name)
	}
}

func (dao *BaseDao[Model]) SavePointTransaction(name string) {
	if dao.transaction != nil {
		dao.transaction.SavePoint(name)
	}
}

func (dao *BaseDao[Model]) WithTransaction(fc func(tx *BaseDao[Model]) error) {

	var callback = func(tx *gorm.DB) error {
		dao.transaction = tx
		defer func() {
			dao.transaction = nil
		}()

		return fc(dao)
	}

	dao.db.Transaction(callback)
}

func (dao *BaseDao[Model]) GetOrm(config ...ORMConfig) *gorm.DB {
	if dao.transaction != nil {
		return dao.transaction
	}
	return dao.db
}

func (dao *BaseDao[Model]) CommitTransaction() {
	if dao.transaction != nil {
		dao.transaction.Commit()
		dao.transaction = nil
	}
}

func (dao *BaseDao[Model]) Create(record Model, config ...ORMConfig) (err error) {
	orm := dao.GetOrm(config...)
	err = orm.Create(&record).Error
	return
}

var ErrorIdIsNotProvided = errors.New("id is not provided")
var ErrorRecordIsNotAStructInstance = errors.New("record is not a struct instance or instance pointer")

func (dao *BaseDao[Model]) UpdateById(record Model, config ...ORMConfig) (err error) {
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

	orm := dao.GetOrm(config...)
	err = orm.Model(model.Elem().Interface()).Updates(record).Error
	return
}

var ErrorNotADbModel = errors.New("the record is not a db model")

func (dao *BaseDao[Model]) LogicalDelete(record Model, table ...ORMConfig) (err error) {

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

	return dao.GetOrm().Model(record).Update("deleted", Deleted).Error

}

func (dao *BaseDao[Model]) FindByKey(key string, value any, table ...ORMConfig) (result *Model, err error) {
	orm := dao.GetOrm()
	if len(table) > 0 {
		orm = orm.Table(table[0])
	}
	err = orm.Where(fmt.Sprintf("%s = ?", key), value).First(result).Error
	return
}
