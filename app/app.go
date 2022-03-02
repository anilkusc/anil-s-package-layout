package app

import (
	"github.com/anilkusc/go-package-layout/api"
	"github.com/anilkusc/go-package-layout/domain"
	"github.com/anilkusc/go-package-layout/pkg"
	"github.com/anilkusc/go-package-layout/repository"
)

type App struct {
}

func (app *App) Init() {
	var err error
	repository := repository.Repository{}
	err = repository.Init()
	if err != nil {
		panic(err)
	}
	packg := pkg.Pkg{
		Repository: &repository,
	}
	err = packg.Init()
	if err != nil {
		panic(err)
	}
	dmn := domain.Domain{
		Pkg: &packg,
	}
	err = dmn.Init()
	if err != nil {
		panic(err)
	}
	api := api.Api{
		Domain: &dmn,
	}
	api.Start()

}

func (app *App) Start() {
	app.Init()
}
