package _package

import (
	"fmt"
	"gostudy.com/pkg_parent"

	//. "fmt" // 取别名为.，则调用的时候可以直接使用包下面的函数
	"gostudy.com/pkg_son"
	stc "strconv" // 为包取别名，使用的时候可以直接使用包的别名
)

func Introduction() {
	fmt.Println("Go语言包管理")
	//Println("Go语言包管理") // 直接使用fmt包下的函数

	var num = 10
	//Println("num=", num, "  转为字符串：", strconv.Itoa(num))
	fmt.Println("num=", num, "  转为字符串：", stc.Itoa(num)) // 使用包别名调用函数

	pks := pkg_son.Params{Num1: 30, Num2: 20}
	sum := pkg_son.Add(pks)
	fmt.Println(pks.Num1, "+", pks.Num2, "=", sum)
	sub := pks.Subtraction()
	fmt.Println(pks.Num1, "-", pks.Num2, "=", sub)

	pkg_son.Son()
	parent := pkg_parent.Data{20, "leedaning"}
	parent.Save()
}
