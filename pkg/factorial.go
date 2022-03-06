package pkg

import (
	"github.com/anilkusc/go-package-layout/repository"
	"gorm.io/gorm"
)

type Factorial struct {
	gorm.Model
	Input  int
	Result int
}

func (f *Factorial) Calculate(repository *repository.Repository) (int, error) {
	f.Result = 1
	for i := 1; i < f.Input+1; i++ {
		f.Result = f.Result * i
	}
	res := repository.Database.Sqlite.Create(f)
	return f.Result, res.Error
}
