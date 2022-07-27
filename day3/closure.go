package day3

import "fmt"

//闭包
func MyClosure() {
	//go语言支持函数式编程：支持将一个函数作为另一个函数的参数，也支持将一个函数作为另一个函数的返回值。
	//一个外层函数中，有内层函数，该内层函数中，会操作外层函数的局部变量（外层函数中的参数，或者外层函数中直接定义的变量），并且该外层函数的返回值就是这个内层函数。这个内层函数和外层函数的局部变量，统称为闭包结构。

	res := increment()
	fmt.Printf("res的类型为：%T \n", res)
	fmt.Println("res的值为：", res)

	v1 := res()
	fmt.Println("v1 = ", v1)
	v2 := res()
	fmt.Println("v2 = ", v2)

	resu := increment()
	fmt.Printf("resu的类型为：%T \n", resu)
	fmt.Println("resu的值为：", resu)

	val1 := resu()
	fmt.Println("val1 = ", val1)
	val2 := resu()
	fmt.Println("val2 = ", val2)

}

func increment() func() int { // 外层函数

	//定义一个局部变量
	i := 0

	//定义一个匿名函数，给变量自增并返回
	fun := func() int { // 内层函数
		i++
		return i
	}

	//返回函数类型的变量
	return fun
}
