package day3

import "fmt"

//回调函数

func MyCallBack() {

	/*fmt.Println("回调函数")
	fmt.Printf("Adds1函数的类型为%T, 地址为：%p\n", Adds1, Adds1)
	fmt.Printf("Adds2函数的类型为%T, 地址为：%p\n", Adds2, Adds2)
	fmt.Printf("Multiplication函数的类型为%T, 地址为：%p\n", Multiplication, Multiplication)
	fmt.Printf("Oper函数的类型为%T, 地址为：%p\n", Oper, Oper)
	fmt.Printf("Oper2函数的类型为%T, 地址为：%p\n", Oper2, Oper2)*/
	//handle()
	num1 := 5
	num2 := 6
	res1 := Oper(num1, num2, Adds1)
	fmt.Println("最终结果为：", res1)
	res2 := Oper2(num1, num2, Adds2)
	fmt.Println("最终结果为：", res2)
	res11 := Oper(num1, num2, Multiplication)
	fmt.Println("最终结果为：", res11)
	res12 := Oper(num1, num2, func(num1, num2 int) int {
		return num2 - num1
	})
	fmt.Println("最终结果为：", res12)
}

func Oper(num1, num2 int, fuc func(int, int) int) int {
	fmt.Println("处理操作的函数 Oper()")
	fmt.Printf("函数变量的类型为：%T \n", fuc)
	fmt.Println("所有参数：", num1, num2, fuc)
	res := fuc(num1, num2)
	return res
}
func Oper2(num1, num2 int, fuc func(...int) int) int {
	fmt.Println("处理操作的函数 Oper()")
	fmt.Printf("函数变量的类型为：%T \n", fuc)
	fmt.Println("所有参数：", num1, num2, fuc)
	res := fuc(num1, num2)
	return res
}

//两个数相加的函数
func Adds1(num1, num2 int) int {
	res := num1 + num2
	return res
}

//两个数相加的函数
func Adds2(num ...int) int {

	var res int
	for _, v := range num {
		res += v
	}
	return res
}

func Multiplication(num1, num2 int) int {
	return num1 * num2
}

func handle() {
	fmt.Println("处理操作的函数 handle()")

	num1 := 10
	num2 := 20
	total := Adds1(num1, num2)
	fmt.Printf("%d + %d = %d\n", num1, num2, total)
}
