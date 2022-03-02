package repository

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	Database *Database
}

func (repository *Repository) Init() error {
	var err error
	repository.Database = &Database{}
	repository.Database.DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
