package pkg

import (
	"testing"
	"time"

	"github.com/anilkusc/go-package-layout/repository"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func Construct() (Pkg, Factorial, error) {

	godotenv.Load("../.env")

	packg := Pkg{
		Repository: &repository.Repository{},
	}

	factorial := Factorial{
		Model: gorm.Model{
			//ID:        1,
			UpdatedAt: time.Time{}, CreatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
		},
		Input:  4,
		Result: 24,
	}
	err := packg.Repository.Init()
	if err != nil {
		return packg, factorial, err
	}
	err = packg.Init()
	if err != nil {
		return packg, factorial, err
	}
	return packg, factorial, nil
}

func Destruct(db *repository.Database) {
	db.Sqlite.Exec("DROP TABLE factorials")
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
		Destruct(repository.Repository.Database)
	}
}

func TestInit(t *testing.T) {
	p, _, err := Construct()
	if err != nil {
		t.Errorf("Error is: %v . Expected: %v", err, nil)
	}
	tests := []struct {
		err error
	}{
		{err: nil},
	}
	for _, test := range tests {
		err := p.Init()
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		Destruct(p.Repository.Database)
	}
}
