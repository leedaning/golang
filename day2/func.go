package day2

import "fmt"

func MyFunc() {
	fmt.Println("函数使用及参数传递")
	var num int = 10
	fmt.Printf("变量num的值为：%d\n", num)
	num = myAdd(num, 2)
	fmt.Printf("变量num的值为：%d\n", num)

	intType, strType, boolType, floatType, intArrType, structType := mySum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("intType = %d, 类型为：%T\n", intType, intType)
	fmt.Printf("strType = %s, 类型为：%T\n", strType, strType)
	fmt.Printf("boolType = %b, 类型为：%T\n", boolType, boolType)
	fmt.Printf("floatType = %f, 类型为：%T\n", floatType, floatType)
	fmt.Printf("intArrType = %d, 类型为：%T\n", intArrType, intArrType)
	fmt.Printf("structType = %d, 类型为：%T\n", structType, structType)
}

func myAdd(num int, step int) int {
	return num + step
}

type mystruct struct {
	name string
	age  int
}

func mySum(num ...int) (int, string, bool, float64, []int, mystruct) {
	var temp int
	for k, v := range num {
		fmt.Println("k=", k, " v=", v)
		temp += v
	}
	fmt.Println("temp=", temp)

	slc := make([]int, 0)
	myStruct := mystruct{"leen", 18}

	return temp, "string", true, 3.14, slc, myStruct
}
