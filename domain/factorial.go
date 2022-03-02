package domain

import "github.com/anilkusc/go-package-layout/pkg"

func (domain *Domain) Calculate() error {
	factorial := pkg.Factorial{}
	factorial.Input = 4
	return factorial.Calculate(domain.Pkg.Repository)
}
