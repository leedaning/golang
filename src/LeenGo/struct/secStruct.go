package main

import (
	. "fmt"
)

type skills []string

type Human struct {
	name  string
	age   int
	wight int
}

type Student struct {
	Human
	speciality string
}

type Teacher struct {
	Human
	skills
	speciality string
}

func main() {
	// firFunc()
	secFunc()
}

func firFunc() {

	mark := Student{Human{name: "leen", age: 30, wight: 60}, "计算机科学与技术"}
	// mark := Student{Human{"Mark", 25, 120}, "Computer Science"}
	mark.Human = Human{name: "leedaing", age: 18, wight: 61}
	mark.Human.age = mark.Human.age - 1

	Println("name:", mark.name)
	Println("age:", mark.age)
	Println("wight:", mark.wight)
	Println("speciality:", mark.speciality)
	Println("mark:", mark.Human)
}

func secFunc() {
	// mark := Teacher{Human{name: "daning", age: 30, wight: 60}, skills{"abc"}, "Biology"}
	mark := Teacher{Human: Human{name: "daning", age: 30, wight: 60}, skills: skills{"abc"}, speciality: "Biology"}

	Println("Teacher info:", mark)
}
