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
	var objects []interface{}
	// get type of the object
	objectType := reflect.TypeOf(object)
	// create new object from obtained Object type. This line is creating new *[]*interface{}. Because gorm is only accept this kind of object on interface array.
	newObject := reflect.New(reflect.SliceOf(objectType)).Interface()
	// list all appropriate objects from the database
	result := d.DB.Model(object).Find(newObject)
	// convert newObject(interface{}) to ([]interface{})
	v := reflect.ValueOf(newObject).Elem()
	// append all of the objects to the interface array
	for i := 0; i < v.Len(); i++ {
		// Getting interface of newObject[i]
		elem := v.Index(i).Interface()
		//append to objects
		objects = append(objects, elem)
	}
	return objects, result.Error
}

func (d *Database) Read(object interface{}) (interface{}, error) {

	objectType := reflect.TypeOf(object)
	newObject := reflect.New(objectType).Interface()
	result := d.DB.Find(&newObject)
	return object, result.Error
}
