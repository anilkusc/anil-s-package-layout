package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	Database *Database
}

func (repository *Repository) Init() error {
	var err error
	repository.Database = &Database{}
	repository.Database.Sqlite, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	return err
}
