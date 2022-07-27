package day3

import "fmt"

//匿名函数

func MyAnonymity() {
	Prt()

	num := 10

	fmt.Printf("num = %d\n", num)
	func() {
		num := 100
		fmt.Println("这是一个匿名函数内部，新声明的num=", num)
	}()
	fmt.Printf("num = %d\n", num)

	func(name string, age, weight, heigh int) int {
		fmt.Printf("\nname :%s, age:%d, weight:%d, heigh:%d \n", name, age, weight, heigh)
		return 0
	}("Leen", 20, 173, 63)

	res := func(num1, num2 int) int {
		fmt.Println("不带括号的匿名函数内部")

		return num1 + num2
	}
	fmt.Printf("\n调用匿名函数本身的结果为：%d", res(1, 2))
}
func Prt() {
	fmt.Println("这是一个函数")
}
