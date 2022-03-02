package repository

import (
	"testing"
)

func TestAutoMigrate(t *testing.T) {
	repository, testtables, err := Construct()

	if err != nil {
		t.Errorf("Error is: %v . Expected: %v", err, nil)
	}
	tests := []struct {
		input interface{}
		err   error
	}{
		{input: testtables, err: nil},
	}
	for _, test := range tests {
		err := repository.Database.AutoMigrate(test.input)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}

	Destruct(repository.Database)
}
func TestWrite(t *testing.T) {
	repository, testtables, err := Construct()
	if err != nil {
		t.Errorf("Error is: %v . Expected: %v", err, nil)
	}
	err = repository.Database.AutoMigrate(testtables)
	if err != nil {
		t.Errorf("Error is: %v . Expected: %v", err, nil)
	}

	tests := []struct {
		input interface{}
		err   error
	}{
		{input: testtables, err: nil},
	}
	for _, test := range tests {
		err := repository.Database.Write(test.input)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}

	Destruct(repository.Database)
}
