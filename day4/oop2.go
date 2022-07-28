package day4

import "fmt"

//人类结构体
type People struct {
	name string
	age  int
}

type Teacher struct {
	People //匿名结构体
	school string
}

func OOP2() {
	var peo People
	peo.name = "小明"
	peo.age = 30

	peo.eat()

	var tech Teacher
	tech.school = "大学"
	tech.name = "小红"
	tech.age = 18
	fmt.Println(tech.name)   // 子类对象可以直接访问父类字段（提升字段）
	fmt.Println(tech.school) // 子类对象访问自己新增的字段属性
	tech.eat()               // 子类调用父类方法
	tech.teaching()          // 子类对象访问子类对象的方法
	//tech.eat()
}

func (m People) eat() {
	fmt.Println("Can eat.", m)
}

func (m Teacher) teaching() {
	fmt.Println("Can teaching", m)
}

/*func (m Teacher) eat() {
	fmt.Println("重写父类的方法")
}*/
