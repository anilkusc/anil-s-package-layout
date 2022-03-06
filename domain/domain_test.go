package domain

import (
	"testing"

	"github.com/anilkusc/go-package-layout/pkg"
	"github.com/anilkusc/go-package-layout/repository"
	"github.com/joho/godotenv"
)

func Construct() (Domain, error) {

	godotenv.Load("../.env")
	domain := Domain{
		Pkg: &pkg.Pkg{
			Repository: &repository.Repository{},
		},
	}
	err := domain.Pkg.Repository.Init()
	if err != nil {
		return domain, err
	}
	err = domain.Pkg.Init()
	if err != nil {
		return domain, err
	}
	err = domain.Pkg.Repository.Init()
	if err != nil {
		return domain, err
	}
	return domain, nil
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
		domain, err := Construct()
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		Destruct(domain.Pkg.Repository.Database)
	}
}

func TestInit(t *testing.T) {
	d, err := Construct()
	if err != nil {
		t.Errorf("Error is: %v . Expected: %v", err, nil)
	}
	tests := []struct {
		err error
	}{
		{err: nil},
	}
	for _, test := range tests {
		err := d.Init()
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		Destruct(d.Pkg.Repository.Database)
	}
}
