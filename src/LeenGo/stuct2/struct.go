package main

import (
	. "fmt"
)

type person struct {
	name string
	age  int
}

func Older(a, b person) (person, int) {
	if a.age > b.age {
		return a, a.age
	}
	return b, a.age
}

func main() {
	var Tom person
	Tom.name = "Tom"
	Tom.age = 18

	Leen := person{"leen", 30}

	John := person{name: "John", age: 20}

	// Jim := new(person)        // 这种用法会提示错误，不知道是怎么回事。cannot use Jim (type *person) as type person in argument to Older
	// Jim.name = "Jim"
	// Jim.age = 22
	Jim := person{name: "Jim", age: 22}

	tl_older, tl_diff := Older(Tom, Leen)
	jj_older, jj_diff := Older(John, Jim)
	Println("\rPersons :", tl_older)
	Println("age is :", tl_diff)
	Println("\rPersons :", jj_older)
	Println("age is :", jj_diff)
}
