package day2

import (
	"fmt"
	"strconv"
)

func MyString() {

	var name string = "Leen"
	fmt.Println("My name is", name, " add = ", &name)
	/*Change(name)
	fmt.Println("My name is", name, " add = ", &name)*/
	/*var str string = "100"
	Convert(str)*/
	var flt float64 = 3.1415926
	fmt.Printf("浮点型变量flt的值为%f, 类型为:%T, 地址为：%p\n", flt, flt, &flt)
	str := strconv.FormatFloat(flt, 'g', 5, 64)
	fmt.Printf("浮点型变量转换为str的值为%s, 类型为:%T, 地址为：%p\n", str, str, &str)

}

func Change(str string) {

	fmt.Println("str = ", str, " add = ", &str)
	str = "New name"
	fmt.Printf("\nstr = %s, add = %p \n", str, &str)
}

//字符串转换其他类型
func Convert(str string) {
	fmt.Printf("转换前的数据为:%s, 类型为:%T, 地址为:%p\n", str, str, &str)
	num, _ := strconv.Atoi(str)
	fmt.Printf("转换后的数据为:%d, 类型为:%T, 地址为:%p\n", num, num, &num)

	str2 := strconv.Itoa(num)
	fmt.Printf("将num转换为字符串为:%s, 类型为:%T, 地址为:%p\n", str2, str2, &str2)
}
