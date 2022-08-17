package day6

import (
	"fmt"
	"strconv"
)

/*type用于类型定义和类型别名
1.类型定义：type 类型名称 Type
2.类型别名：type 类型名称 = Type*/

//定义结构体
type strt struct {
	name string
}

//定义接口
type intf interface {
	life()
}

//定义myint类型
type myint int

//定义mystr类型
type mystr string

//定义函数类型
type myfun func(int, int) string

func MyType() {
	var i1 myint
	var i2 = 100
	i1 = 200
	fmt.Println("i1 = ", i1, " i2 = ", i2) // i1 =  200  i2 =  100
	fmt.Printf("%T, %T\n", i1, i2)         // day6.myint, int

	//i1 = i2 //错误，类型不匹配（cannot use i2 (variable of type int) as type myint in assignment）
	fmt.Println("--------")
	res1 := fun1()
	fmt.Println(res1(10, 20))
}

//声明一个返回值类型为myfun类型的函数
func fun1() myfun {
	fun := func(a, b int) string {
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}

	return fun
}
