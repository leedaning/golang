package day3

import "fmt"

//函数的类型
func MyFuncType() {

	var fc func(int, int) // 声明一个函数类型变量
	fc = myFunc
	fmt.Println(fc)
	myFunc(1, 2)
	fc(10, 20)
	MyVarType()

}

type myStruct struct {
	name string
	age  int
}

//变量的类型
func MyVarType() {
	var num int = 10
	fmt.Printf("\nnum的类型为：%T", num)

	var str string = "This is string"
	fmt.Printf("\nstr的类型为：%T", str)

	arr := [2]int{1, 2}
	fmt.Printf("\narr的类型为：%T", arr)

	slc := []int{}
	fmt.Printf("\nslc的类型为：%T", slc)

	myMap := make(map[string]string, 0)
	fmt.Printf("\nmyMap的类型为：%T", myMap)

	myStruct := myStruct{name: "leen", age: 20}
	fmt.Printf("\nmyStruct的类型为：%T", myStruct)

	fuc := myFunc // 函数类型
	fmt.Printf("\nfuc的类型为：%T", fuc)

	fuc2 := myFunc2 // 函数类型
	fmt.Printf("\nfuc2的类型为：%T", fuc2)

}

func myFunc(num1, num2 int) {
	fmt.Println("函数参数相加的结果为：", num1+num2)
	//return num1 + num2
}
func myFunc2(name, nick_name string, age, heigh, weight int, sex bool) (string, error) {
	fmt.Println("函数2")
	return name, nil
}
