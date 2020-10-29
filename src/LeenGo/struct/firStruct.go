package main

import (
	. "fmt"
)

// type myFunc func(num int) bool
type person struct {
	name string
	age  int
}

func main() {
	// 第一种
	// var P person
	// P.name = "leen"
	// P.age = 30

	// 第二种
	// P := person{"leen", 30}

	// 第三种
	// P := person{name: "leen", age: 30}

	// 第四种
	P := new(person)
	P.name = "leen"
	P.age = 30

	Println("Hello, this is struct.")
	Printf("\rMy name is %s, age %d", P.name, P.age)
}
