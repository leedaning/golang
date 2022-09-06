package day7

import (
	"errors"
	"fmt"
	"os"
)

//自定义的错误类型
type CustomError struct {
	msg  string // 错误信息
	time string // 错误发生的时间
}

//错误类型
func MyErrorType() {
	//getType()
	/*cust := CustomError{"时间错误", "错误时间"}
	err := cust.LogicError(30, 20)
	if ins, ok := err.(*CustomError); ok {
		fmt.Println("结果为：", ins)
	}*/
	cont, err := getType()
	if err != nil {
		fmt.Println("内容：", cont)
		if ins, ok := err.(*CustomError); ok { //断言错误类型
			fmt.Println("ins=", ins)
			fmt.Println("msg:", ins.msg)
			fmt.Println("time:", ins.time)
			err := ins.LogicError(10, 9)
			if err != nil {
				fmt.Println(err.Error()) // 打印错误信息
			}
		}
	}

}

//获取错误类型
func getType() (string, error) {

	cont, err := os.Open("./day7/test.txt0")
	//cust := CustomError{"文件路径错误", "错误时间"}
	if err != nil {
		fmt.Println("打开文件错误,", err)
		if ins, ok := err.(*os.PathError); ok { // 断言错误类型
			fmt.Println("1. Op:", ins.Op)
			fmt.Println("2. Path:", ins.Path)
			fmt.Println("3. Err:", ins.Err)
		}
		//return "文件没打开", cust.LogicError(10, 2)
		return "文件没打开", &CustomError{"文件路径错误", "错误时间"}
	} else {
		fmt.Println("打开文件成功", cont)
		//return "文件内容", cust.LogicError(2, 1)
		return "文件打开成功", &CustomError{"成功", "打开时间"}
	}
}

//实现逻辑错误方法
func (r *CustomError) LogicError(a, b int) error {
	if a > b {
		return errors.New("逻辑错误哈~")
	} else {
		fmt.Println("a=", a, "b=", b)
	}
	return nil
}

//实现Error接口
func (r *CustomError) Error() string {
	return "错误啦"
}
