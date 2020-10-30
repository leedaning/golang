package main

import (
	. "fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	loan float32
}

type Employee struct {
	Human
	money float32
}

func (h Human) SayHi() {
	Printf("\r Hi my name is %s , say hi!", h.name)
}

func (h Human) Sing(layric string) {
	Printf("\r My name is %s , I will sing is %s", h.name, layric)
}

func (s Student) SayHi() {
	Printf(" Hi i'm a student, name is %s, loan %d", s.name, s.loan)
}

func (e Employee) SayHi() {
	Printf(" Hi i'm a employee, name is %s, money %d", e.name, e.money)
}

type Men interface {
	SayHi()
	Sing(layric string)
}

func main() {
	human := Human{"leen", 30, "17737131001"}
	stu := Student{Human{"daning", 18, "17737131002"}, 100.00}
	emp := Employee{Human{"leedaning", 22, "17737131003"}, 10000.00}

	var m Men

	m = human
	m.SayHi()
	m.Sing("My love")

	m = stu
	m.SayHi()
	m.Sing("I am student")

	m = emp
	m.SayHi()
	m.Sing("I am employee")
}
