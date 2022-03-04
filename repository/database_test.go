package repository

import (
	"fmt"
	"reflect"
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
	repository, testtable, err := Construct()
	if err != nil {
		t.Errorf("Error is: %v . Expected: %v", err, nil)
	}

	tests := []struct {
		input interface{}
		err   error
	}{
		{input: testtable, err: nil},
	}
	for _, test := range tests {
		err := repository.Database.Write(test.input)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}

	Destruct(repository.Database)
}
func TestList(t *testing.T) {
	repository, testtable, err := Construct()
	if err != nil {
		t.Errorf("Error is: %v . Expected: %v", err, nil)
	}
	err = repository.Database.Write(testtable)
	if err != nil {
		t.Errorf("Error is: %v . Expected: %v", err, nil)
	}
	tests := []struct {
		input  interface{}
		output []interface{}
		err    error
	}{
		{input: testtable, output: []interface{}{testtable}, err: nil},
	}
	for _, test := range tests {
		output, err := repository.Database.List(test.input)
		fmt.Println(reflect.TypeOf(output))
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		if !reflect.DeepEqual(test.output, output) {
			t.Errorf("Result is: %v . Expected: %v", output, test.output)
		}
	}

	Destruct(repository.Database)
}
