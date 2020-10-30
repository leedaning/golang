package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

//Human对象实现Sayhi方法
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human对象实现Sing方法
func (h *Human) Sing(lyrics string) {
	fmt.Println("La la la la la la la ", lyrics)
}

//Human对象实现Guzzle方法
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// Employee重载Human的Sayhi方法
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //此句可以分成多行
}

//Student实现BorrowMoney方法
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and...)
}

//Employee实现SpendSalary方法
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!! Get me through the day!
}

type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(BeerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(lyrics string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(lyrics string)
	SpendSalary(amount float32)
}

// func main() {
// 	fmt.Println("Hello")

// 	// stu := Student:Student{Human:Human{name:"leen", age:18, phone:"17737131001"}, school:"First school", loan:10.00}
// 	stu := Student{Human{"leen", 18, "17737131001"}, "First school", 10.00}

// 	var intf YoungChap
// 	intf = stu
// 	intf.SayHi()
// 	intf.Sing("Let's go")
// 	intf.BorrowMoney(10.11)
// }
