package domain

import (
	"errors"

	"github.com/anilkusc/go-package-layout/pkg"
)

func (domain *Domain) Calculate(input int) (int, error) {
	if input < 1 {
		return 0, errors.New("number cannot be smaller than 1")
	}
	factorial := pkg.Factorial{}
	factorial.Input = input
	return factorial.Calculate(domain.Pkg.Repository)
}
