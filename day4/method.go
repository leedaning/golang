package day4

import "fmt"

type Car struct {
	name       string // 名称
	engine     string // 发动机
	createTime string // 制造时间
}

//油动力车
type OilCar struct {
	fuel string
	name string
	car  Car
}

//新能源车
type NewEnergy struct {
	energy string
}

//方法
func Method() {
	car := Car{
		name:       "汽车",
		engine:     "蒸汽",
		createTime: "2000",
	}
	var carP *Car
	carP = &car
	oilCar := OilCar{
		"汽油",
		"汽油车",
		Car{
			"汽车",
			"蒸汽机",
			"1900",
		},
	}

	car.Car()
	carP.Car()
	car.Car2()
	carP.Car2()
	oilCar.Car()

}

//汽车方法(指针）
func (m *Car) Car() {
	fmt.Println("汽车方法", m)

}

//汽车方法（值）
func (m Car) Car2() {

	fmt.Println("汽车方法2", m)
}

//燃油车方法（指针），方法可以同名，只要接收者不同即可
func (m *OilCar) Car() {

	fmt.Println("燃油车", m)
}
