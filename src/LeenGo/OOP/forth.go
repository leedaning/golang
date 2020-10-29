package main

import (
	"fmt"
)

type Human struct {
	name   string
	age    int
	iphone string
}

type Student struct {
	Human
	school string
}

type Teacher struct {
	Human
	school string
	course string
}

func (h Human) SayHi() {
	fmt.Printf("Hi, I am is %s you can call me on %s \n", h.name, h.iphone)
}

func (s *Student) SayHi() {
	fmt.Printf("Hi, I am is %s you can call me on %s, my school is %s \n", s.name, s.iphone, s.school)
}

func (t *Teacher) SayHi() {
	fmt.Printf("Hi, I am is %s you can call me on %s, my school is %s, course is %s \n", t.name, t.iphone, t.school, t.course)
}

func main() {
	stu := Student{Human{"小明", 18, "17737131001"}, "第一中学"}
	ter := Teacher{Human{"李杰", 30, "17737131002"}, "第一中学", "Computer Science"}

	fmt.Println("The student info is ", stu)
	fmt.Println("The teacher info is ", ter)

	stu.SayHi()
	ter.SayHi()
}
