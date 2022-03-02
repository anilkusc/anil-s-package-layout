package domain

import "github.com/anilkusc/go-package-layout/pkg"

type Domain struct {
	Pkg *pkg.Pkg
}

func (domain *Domain) Init() error {
	return nil
}
