package day5

import "fmt"

//空接口，没有包含任何方法的接口
type Null interface {
}

//结构体类型
type Student struct {
	id   int
	name string
}

// 执行函数
func MyNull() {
	var i int = 100
	var s string = "name"
	var f float64 = 3.14
	var stct Student = Student{1, "Leen"}
	fmt.Printf("变量i的值为:%d, 类型为：%T, 内存地址为：%p\n", i, i, &i)
	fmt.Printf("变量s的值为:%s, 类型为：%T, 内存地址为：%p\n", s, s, &s)
	fmt.Printf("变量f的值为:%f, 类型为：%T, 内存地址为：%p\n", f, f, &f)
	fmt.Printf("变量stct的值为:%d, 类型为：%T, 内存地址为：%p\n", stct, stct, &stct)
	/*scan(i)
	scan(s)
	scan(f)
	scan(stct)*/
	/*scan2(i)
	scan2(s)
	scan2(f)
	scan2(stct)*/
	scans2(i, s, f, stct)
}

//获取输入参数的函数，空接口作为参数可以传入任何类型的参数
func scan(p Null) {
	fmt.Printf("参数类型为:%T\n", p)
}

//也可以使用匿名接口作为参数
func scan2(p interface{}) {
	fmt.Printf("参数类型为:%T\n", p)
}

//以空接口作为参数的多参数函数
func scans(p ...Null) {
	for k, v := range p {
		fmt.Println("k=", k, " v=", v)
	}
}

//以匿名空接口作为参数的多参数函数
func scans2(p ...interface{}) {
	for k, v := range p {
		fmt.Println("k=", k, " v=", v)
	}
}
