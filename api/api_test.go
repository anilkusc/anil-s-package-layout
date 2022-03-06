package api

import (
	"testing"

	"github.com/anilkusc/go-package-layout/domain"
	"github.com/anilkusc/go-package-layout/pkg"
	"github.com/anilkusc/go-package-layout/repository"
	"github.com/joho/godotenv"
)

func Construct() (Api, error) {

	godotenv.Load("../.env")
	api := Api{
		Domain: &domain.Domain{
			Pkg: &pkg.Pkg{
				Repository: &repository.Repository{},
			}}}

	err := api.Domain.Pkg.Repository.Init()
	if err != nil {
		return api, err
	}
	err = api.Domain.Pkg.Init()
	if err != nil {
		return api, err
	}
	err = api.Domain.Init()
	if err != nil {
		return api, err
	}

	api.Init()

	return api, nil
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
		api, err := Construct()
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		Destruct(api.Domain.Pkg.Repository.Database)
	}
}
