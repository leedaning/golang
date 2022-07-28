package day4

import "fmt"

//人类结构体
type Person struct {
	name string
	sex  string
	age  int
}

//工人结构体
type Worker struct {
	company string
	Person  // 匿名结构体，相当于可以直接使用woker.name等Person中的字段
}

//学生结构体
type Student struct {
	school string
	Person // 匿名结构体，相当于可以直接使用woker.name等Person中的字段
}

func OOP() {
	worker()
}

func worker() {
	per := Person{
		"Leen",
		"male",
		22,
	}
	worker := Worker{
		"Internet",
		per,
	}
	student := Student{}
	student.school = "University"
	student.Person = per
	fmt.Println("人类信息：", per)
	fmt.Println("工人信息：", worker)
	fmt.Println("学生信息：", student)
	student.name = "John"
	fmt.Println("学生信息：", student)
}
