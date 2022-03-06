package repository

import (
	"testing"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func Construct() (Repository, interface{}, error) {

	godotenv.Load("../.env")
	repository := Repository{}
	type TestingPurposeStruct struct {
		gorm.Model
		Name string
		Role string
	}

	var tst = TestingPurposeStruct{
		Model: gorm.Model{
			//ID:        1,
			UpdatedAt: time.Time{}, CreatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
		},
		//Test: 1,
		Name: "test",
		Role: "admin",
	}
	err := repository.Init()
	if err != nil {
		return repository, tst, err
	}
	err = repository.Database.Sqlite.AutoMigrate(&TestingPurposeStruct{})
	if err != nil {
		return repository, tst, nil
	}
	return repository, tst, nil
}

func Destruct(db *Database) {
	db.Sqlite.Exec("DROP TABLE testing_purpose_structs")
}

func TestConstruct(t *testing.T) {

	tests := []struct {
		err error
	}{
		{err: nil},
	}
	for _, test := range tests {
		repository, _, err := Construct()
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		Destruct(repository.Database)
	}
}
func TestInit(t *testing.T) {
	repository, _, err := Construct()
	if err != nil {
		t.Errorf("Error is: %v . Expected: %v", err, nil)
	}
	tests := []struct {
		err error
	}{
		{err: nil},
	}
	for _, test := range tests {
		err := repository.Init()
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		Destruct(repository.Database)
	}
}
