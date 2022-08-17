package day5

import "fmt"

//动物接口
type Animal interface {
	live()
	think()
}

//鸟类接口
type Birds interface {
	Animal // 接口嵌套
	fly()
}

//鱼类接口
type Fishs interface {
	swiming()
}

//鸟类结构体
type Bird struct {
	name string
}

//鱼类结构体
type Fish struct {
	name string
}

func MyItfe() {
	var phoenix = Bird{"凤凰"} // 实现了Animal接口和Birds接口
	MyLoad(phoenix)
	phoenix.fly()

	shark := Fish{"鲨鱼"} // 实现了Animal接口和Fishs接口
	MyLoad(shark)
	shark.swiming()

	//接口类型的变量
	var anml Animal
	anml = phoenix // 子类对象赋给父类对象，父类对象只能访问自己的函数和变量，不能访问子类的对象和方法，如可以访问live()、think()，不能访问name属性和fly()函数
	//anml.live()
	MyLoad(anml)

}

//一个函数如果接受接口类型作为参数，那么实际上可以传入该接口的任意实现类型对象作为参数
//定义一个类型为接口类型，实际上可以赋值为任意实现类的对象
//调用接口共有方法
func MyLoad(A Animal) {
	A.live()
	A.think()
}

//Bird结构体实现了live()、think()、fly()，则即实现了Animal接口又实现了Birds接口
func (r Bird) live() {
	fmt.Println(r.name + "有生命")
}

func (r Bird) think() {
	fmt.Println(r.name + "会思考")
}

func (r Bird) fly() {
	fmt.Println(r.name + "会飞翔")
}

func (r Fish) live() {
	fmt.Println(r.name + "有生命")
}

func (r Fish) think() {
	fmt.Println(r.name + "会思考")
}

func (r Fish) swiming() {
	fmt.Println(r.name + "会游泳")
}
