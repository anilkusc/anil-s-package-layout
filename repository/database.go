package repository

import (
	"gorm.io/gorm"
)

type Database struct {
	Sqlite *gorm.DB
}
