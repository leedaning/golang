package day7

import (
	"errors"
	"fmt"
	"os"
)

//错误处理学习
func MyError() {
	//fmt.Println("错误处理学习")
	//study()
	res := checkAge(-18)
	if res != nil {
		fmt.Println("错误信息为:", res)
	}

}

func study() {
	//error:内置的数据类型，内置的接口
	//定义方法：Error() string

	//使用go语言提供好的包：
	//errors包下的函数，New(),创建一个error(错误)对象
	//fmt包下的函数，Errorf()，func Errorf(format string, a ...any) error

	cont, err := os.Open("./day7/test.txt0")
	if err != nil {
		//log.Fatal(err)
		fmt.Println("打开文件错误,", err)
	} else {
		fmt.Println("打开文件成功", cont)
	}

	err2 := errors.New("自定义的错误")
	fmt.Println(err2)                 //自定义的错误
	fmt.Printf("err2的类型为：%T\n", err2) // err2的类型为：*errors.errorString

	err3 := fmt.Errorf("错误信息哈%s", err2)
	fmt.Println(err3)
	fmt.Printf("err3的类型为：%T\n", err3) // err3的类型为：*errors.errorString
}

//设计一个函数，也可以返回error（错误信息）。验证年龄，如果年龄不合法就返回错误信息。
func checkAge(age int) error {
	if age < 0 {
		//return errors.New("年龄不合法")
		return fmt.Errorf("年龄%d不合法", age)
	}
	fmt.Println("年龄是：", age)
	return nil
}
