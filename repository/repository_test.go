package repository

import (
	"testing"

	"github.com/joho/godotenv"
)

func Construct() (Repository, interface{}, error) {

	godotenv.Load("../.env")
	repository := Repository{}
	type TestingPurposeStruct struct {
		Name string
		Role string
	}

	tst := TestingPurposeStruct{
		Name: "test",
		Role: "admin",
	}
	err := repository.Init()
	if err != nil {
		return repository, tst, err
	}
	err = repository.Database.AutoMigrate(&TestingPurposeStruct{})
	if err != nil {
		return repository, tst, nil
	}
	return repository, tst, nil
}

func Destruct(db *Database) {
	db.DB.Exec("DROP TABLE testing_purpose_structs")
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
