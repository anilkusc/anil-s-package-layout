package repository

import (
	"reflect"

	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func (d *Database) AutoMigrate(object interface{}) error {
	return d.DB.AutoMigrate(&object)
}

func (d *Database) Write(object interface{}) error {
	result := d.DB.Create(object)
	return result.Error
}
func (d *Database) List(object interface{}) ([]interface{}, error) {
	objects := []interface{}{}
	dtype := reflect.TypeOf(object)
	pages := reflect.New(dtype).Interface()

	result := d.DB.Find(&pages) // d.DB.Find(reflect.TypeOf(object).Kind())
	//result := d.DB.Model(reflect.New(reflect.SliceOf(reflect.TypeOf(object))).Interface()).Find(&objects) // d.DB.Find(reflect.TypeOf(object).Kind())

	return objects, result.Error
}
