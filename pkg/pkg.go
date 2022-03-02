package pkg

import "github.com/anilkusc/go-package-layout/repository"

type Pkg struct {
	Repository *repository.Repository
}

func (pkg *Pkg) Init() error {
	return pkg.Repository.Database.AutoMigrate(&Factorial{})
}
