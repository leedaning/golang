package intf

import (
	"fmt"
	"reflect"
)

type person struct {
	name string
	age  int
}

type live interface {
	life()
}

func Learn() {
	str()
	fmt.Println("")
	integer()
	fmt.Println("")
	stru()
	fmt.Println("")
	slc()
	fmt.Println("")
	myNil()
}

//字符串类型
func str() {
	var str string
	//str = "今天天气真好"

	fmt.Printf("字符串 str: %s\n", str)
	if str == "" {
		fmt.Print("为空")
	} else {
		fmt.Print("不为空")
	}
	fmt.Println("")
	eptIntf(str)
}

func integer() {
	var num int

	fmt.Printf("整数 num: %s\n", num)
	if num == 0 {
		fmt.Print("为空")
	} else {
		fmt.Print("不为空")
	}
	fmt.Println("")
	eptIntf(num)
}

//结构体类型
func stru() {
	var per person
	/*per.name = "Leen"
	per.age = 30*/
	fmt.Println("结构体per：", per)
	if per == (person{}) {
		fmt.Print("为空")
	} else {
		fmt.Print("不为空")
	}
	fmt.Println("")
	eptIntf(per)
}

func slc() {
	var slc []int
	fmt.Println("切片 slc：", slc)
	if slc == nil {
		fmt.Print("为空")
	} else {
		fmt.Print("不为空")
	}
	fmt.Println("")
	eptIntf(slc)
}

func myNil() {

	fmt.Println("空类型：", nil)
	fmt.Println("")
	eptIntf(nil)
}

func eptIntf(i interface{}) {
	defer func() {
		recover()
	}()
	fmt.Println("空接口类型作为参数", i)
	if i != nil {
		if reflect.ValueOf(i).IsNil() {
			fmt.Println("真的是空的")
		} else {
			fmt.Println("真的不为空")
		}
	} else {
		fmt.Println("空接口参数值为空")
	}

	types := reflect.TypeOf(i) // 获取类型
	val := reflect.ValueOf(i)  // 获取值
	fmt.Println("参数类型：", types)
	fmt.Println("参数类型的类型：", val.Kind()) // 如：intf.person类型的结构体，他的类型就是结构体
	fmt.Println("参数类型：", val.String())
	fmt.Println("是否指针：", reflect.Ptr)
	fmt.Println("是否可以寻找地址：", val.CanAddr()) // 判断是否可以寻找到内存地址
	fmt.Println("是否可以修改：", val.CanSet())    // 可以正确判断一个值是否可以修改。 反射可以读取未导出结构字段的值，但是不能更新这些值。一个可寻址的reflect.Value会记录它是否是通过遍历一个未导出字段来获得的，如果是则不允许修改。所以在更新前使用CanAddr()判断并不保险。CanSet()可以正确判断一个值是否可以修改
	fmt.Println("是否为空：", val.IsNil())
}
