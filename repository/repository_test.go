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
	}
	tst := TestingPurposeStruct{
		Name: "test",
	}
	err := repository.Init()
	if err != nil {
		return repository, nil, err
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
