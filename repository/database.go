package repository

import (
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
