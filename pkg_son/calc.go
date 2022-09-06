package pkg_son

import (
	"gostudy.com/pkg_parent"
)

type Params struct {
	Num1 int
	Num2 int
}

func Son() {
	var data pkg_parent.Data
	data.Id = 10
	data.Name = "Leen"
	data.Save()
}

func Add(par Params) int {
	return par.Num1 + par.Num2
}

func (r Params) Subtraction() int {
	return r.Num1 - r.Num2
}
