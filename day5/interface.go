package day5

import "fmt"

//汽车接口
type Car interface {
	run()
	stop()
}

//汽车结构体
type StctCar struct {
	name string //名称
}

//实现类，巴士车结构体
type StctBus struct {
	//StctCar
	using string // 用途
	name  string
}

//实现类，消防车结构体
type StctFireTruct struct {
	StctCar
	exitinguishing string // 灭火
	name           string
}

//接口
func MyInterface() {
	fmt.Println("接口")

	car := StctCar{name: "第一部小汽车"}
	car.run()

	bus := StctBus{using: "载客", name: "载客巴士车"}
	/*var bus StctBus
	bus.name = "巴士"
	bus.using = "载客"*/
	/*bus.run()
	bus.stop()*/
	load(bus)
	//bus.name  // 子类可以调用自己声明的属性
	//bus.use() // 子类可以调用自己声明的函数

	var myCar Car
	myCar = bus
	//myCar.name  // 父类不能调用子类的属性
	//myCar.use() // 父类不能调用子类的函数
	load(myCar)

	/*var fireTruct StctFireTruct
	//fireTruct := StctFireTruct{name: "消防车", exitinguishing: "灭火"}
	fireTruct.name = "消防车"
	fireTruct.exitinguishing = "灭火"
	fireTruct.run()
	fireTruct.extinguish()*/

}

func load(car Car) {
	fmt.Println("load函数")
	car.run()
	car.stop()
}

//汽车run()方法
func (m StctCar) run() {
	fmt.Println("汽车，可以跑起来的：", m.name)
}

//汽车stop()方法
func (m StctCar) stop() {
	fmt.Println("汽车，可以停下来的：", m.name)
}

//巴士车run()方法，不可以使用(m *StctBus)，不然load中无法使用
func (m StctBus) run() {
	fmt.Println("巴士车，可以跑起来的：", m.name)
}

//巴士车stop()方法
func (m StctBus) stop() {
	fmt.Println("巴士车，可以停下来的：", m.name)
}

//巴士车用途方法use()，实现类自己声明的函数
func (m StctBus) use() {
	fmt.Println("可以载客")
}

//消防车方法run()
func (m StctFireTruct) run() {
	fmt.Println("消防车，可以跑起来的：", m.name)
}

func (m StctFireTruct) extinguish() {
	fmt.Println("可以灭火")
}
