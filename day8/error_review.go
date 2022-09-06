package day8

import (
	"errors"
	"fmt"
	"os"
)

//定义错误类型结构体CustomizeError
type CustomizeError struct {
	code int    // 错误代码
	msg  string // 错误信息
}

func MyError() {
	//error01()
	//error02()
	//getErrorType()
	res, err := Customize(200) // 调用自定义函数，自定义函数中会有返回自定义错误类型
	if err != nil {
		fmt.Println("有错误")

		if ins, ok := err.(*CustomizeError); ok {
			fmt.Println("错误代码：", ins.code, "|| 错误信息：", ins.msg)

			str := ins.RangeError()
			fmt.Println("结果信息：", str)

		}
	}
	fmt.Println(res)
}

//初识错误
func error01() {

	f, err := os.Open("./README.md")
	if err != nil {
		fmt.Println("文件打开失败:", err)
	} else {
		fmt.Println("文件", f.Name(), "打开成功，文件内容为：")
	}
}

//新建错误的两种方法
func error02() {
	//方法一：使用errors.New()
	err := errors.New("新建的错误")
	fmt.Println(err)

	//方法二：使用fmt.Errorf()
	err2 := fmt.Errorf("这也是新建的错误")
	fmt.Println(err2)
}

//获取错误的类型
func getErrorType() {
	f, err := os.Open("../README.md")
	if err != nil {
		if ins, ok := err.(*os.PathError); ok {
			fmt.Println("Error != nil")
			fmt.Println("Error Op is :", ins.Op)
			fmt.Println("Error path", ins.Path)
			fmt.Println("Error is :", ins.Err)
		}
		return
	} else {
		fmt.Println("文件", f.Name(), "打开成功，文件内容为：")
	}
}

//结构体CustomizeError实现error的Error()接口
func (r CustomizeError) Error() string {
	return fmt.Sprintf("code:", r.code, "message:", r.msg)
}

//自定义错误方法：RangeError
func (r CustomizeError) RangeError() string {
	return fmt.Sprintf("code:", r.code, "message", r.msg)
}

// 自定义函数，返回自定义错误类型
func Customize(age int) (string, error) {
	if age >= 150 {
		return "年龄值超出范围", &CustomizeError{10001, "年龄不能大于150"}
	}

	if age < 0 {
		return "年龄不合法", &CustomizeError{code: 1000, msg: "年龄不能小于零"}
	}

	return "年龄合法", nil
}
