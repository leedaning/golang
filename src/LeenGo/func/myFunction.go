package main

import (
	"fmt"
)

func main() {
	var name string = "leen"
	var age int = 30
	myAge, myName := myFunc(name, age)
	fmt.Printf("\r\rMy name is %s, my age is %d", myName, myAge)

	x := 3
	y := 4

	exchangeVal(x, y)
	res := increment(x)
	fmt.Printf("\r res = %d, x = %d", res, x)
	resu := increment2(&x)
	fmt.Printf("\r resu = %d, x = %d", resu, x)
	// changeParam(10, 20)
	// myPainc()

}

/*
 数字变量交换
*/
func exchangeVal(x, y int) {

	// 第一种
	// xPLUSy, xTIMESy := sumAndProduct(x, y)
	// fmt.Printf("\r %d + %d = %d", x, y, xPLUSy)
	// fmt.Printf("\r %d * %d = %d", x, y, xTIMESy)

	// 第二种，巧妙利用异或运算符，而且不用额外声明变量
	// x = x ^ y
	// y = x ^ y
	// x = x ^ y
	// fmt.Printf("\r x = %d, y = %d", x, y)

	// 第三种，使用中转变量
	// var temp int
	// temp = x
	// x = y
	// y = temp

	fmt.Printf("\r\r x = %d, y = %d", x, y)
}

func sumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

func myFunc(name string, age int) (myAge int, myName string) {
	return age, name
}

/*
可变参数，参数数量不确定
*/
func changeParam(arg ...int) {
	for _, val := range arg {
		fmt.Printf("\r val = %d", val)
	}
}

/*
参数值传递
*/
func increment(num int) (res int) {
	res = num + 1
	return
}

/*
参数址传递，引用传递
*/
func increment2(num *int) int {
	*num = *num + 1
	return *num
}

func myPainc() {
	var num int = 100
	defer func() {
		fmt.Println("\r\rHaha")
	}()
	fmt.Println("\r Num is ", num)
	panic("\rHave exception")
	num++
	fmt.Println("\r Num is ", num)
}
