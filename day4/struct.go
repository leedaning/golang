package day4

import "fmt"

type person struct {
	name string
	sex  string
	age  int
}

type student struct {
	//person // person结构体类型，匿名字段
	Person person
	string // 班级，匿名字段
	//string         // 学校
	Student_NO int // 学号
}

//结构体
func MyStruct() {

	//perFuc()
	studentFuc()

	/*strct := struct {
		name string
		sex  string
		age  int
	}{
		"Leen",
		"male",
		22,
	}
	fmt.Println("匿名结构体为：", strct)*/

}

//学生函数
func studentFuc() {
	var std student
	//var per person
	per := person{
		name: "Leen",
		sex:  "male",
		age:  22,
	}
	/*std.Person.name = "Leen"
	std.Person.sex = "male"
	std.Person.age = 22*/
	std.Person = per
	std.string = "二班"
	std.Student_NO = 1
	fmt.Println("人物信息：", per)
	fmt.Println("学生信息：", std)
	fmt.Println("学生信息中人物信息：", std.Person)
	//per.age = 23
	std.Person.age = 24
	fmt.Println("人物信息：", per)
	fmt.Println("学生信息：", std)
	fmt.Println("学生信息中人物信息：", std.Person)
}

//人物函数
func perFuc() {
	//var person person
	/*person := person{}

	person.name = "Leen"
	person.sex = "male"
	person.age = 20*/

	//person := person{name: "leen", sex: "male", age: 22}

	person := person{"leen", "male", 22}

	fmt.Println("人物信息：", person)
	fmt.Println("name:", person.name)
}
