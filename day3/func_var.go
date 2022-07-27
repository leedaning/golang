package day3

import (
	"fmt"
)

//函数参数类型
func FuncVar() {

	//函数的变量类型与变量的类型类似，分基本数据类型和复合数据类型
	//基本数据类型有：
	//	int float bool string
	//复合数据类型有：
	//	array slice map channel struct func interface
	base(100, 3.14, true, "leen")

	fmt.Println()

	var arr = [2]int{10, 20}
	fmt.Printf("数组arr：%d, Type is %T, add is %p\n", arr, arr, &arr)
	var slc = []int{100, 200}
	fmt.Printf("切片slc：%d, Type is %T, add is %p\n", slc, slc, &slc)

	resArr, resSlc := complex(arr, slc)
	fmt.Printf("数组参数返回结果：%d, Type is %T, add is %p\n", resArr, resArr, &resArr) // 数组是值传递
	fmt.Printf("切片slc：%d, Type is %T, add is %p\n", slc, slc, &slc)             // 切片是址传递
	fmt.Printf("切片参数返回结果：%d, Type is %T, add is %p\n", resSlc, resSlc, &resSlc)
}

//基本数据类型作为函数参数
func base(num int, flot float64, bl bool, str string) int {
	fmt.Printf("num = %d, Type is %T\n", num, num)
	fmt.Printf("flot = %f, Type is %T\n", flot, flot)
	fmt.Printf("num = %b, Type is %T\n", num, num)
	fmt.Printf("num = %s, Type is %T\n", num, num)
	return 0
}

//复合类型变量作为函数参数
func complex(arr [2]int, slc []int) ([2]int, []int) {

	fmt.Printf("\n数组类型的参数， arr:%d, type = %T, add = %p \n", arr, arr, &arr)
	for k, v := range arr {
		fmt.Println("k = ", k, " v = ", v)
	}
	arr[0] = 101

	fmt.Printf("\nslice类型的参数，slc:%d, type = %T, add = %p \n", slc, slc, &slc)
	for k, v := range slc {
		fmt.Println("k = ", k, " v = ", v)
	}
	slc[0] = 202

	return arr, slc
}
