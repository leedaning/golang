package main

import (
	"fmt"
)

type Human struct {
	name string
	age  int
}

type Student struct {
	Human
	speciality string
}

type Teacher struct {
	Human
	speciality string
}

func main() {
	Jim := Student{Human: Human{name: "Jim", age: 18}, speciality: "technology"}
	Tom := Teacher{Human: Human{name: "Tom", age: 30}, speciality: "Computer Science"}

	fmt.Println(Jim)
	fmt.Println(Tom)
}
