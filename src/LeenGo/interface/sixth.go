package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name string
	age  int
}

type Element interface{}
type List []Element

func (h Human) String() string {
	return "(name" + h.name + ", age : " + strconv.Itoa(h.age) + " year"
}

func main() {
	list := make(List, 4)
	list[0] = "leen"
	list[1] = 18
	list[2] = [2]string{"leen", "male"}
	list[3] = Human{"leedaning", 30}

	// getType(list)
	getType2(list)

}

/*
根据interface变量反推出变量类型
方法一：
Go语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。
如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false。
*/
func getType(list List) {

	for k, v := range list {
		if value, ok := v.(int); ok {
			fmt.Printf("\r list[%d] is a int and its value is %v", k, value)
		} else if value, ok := v.(string); ok {
			fmt.Printf("\r list[%d] is a string and its value is %v", k, value)
		} else if value, ok := v.(Human); ok {
			fmt.Printf("\r list[%d] is a Human and its value is %v", k, value)
		} else {
			fmt.Printf("\r list[%d] is a different type", k)
		}
	}
}

func getType2(list List) {

	for k, v := range list {
		switch value := v.(type) {
		case int:
			fmt.Printf("\r list[%d] is a int and its value is %v", k, value)
		case string:
			fmt.Printf("\r list[%d] is a string and its value is %v", k, value)
		case Human:
			fmt.Printf("\r list[%d] is a Human and its value is %v", k, value)
		default:
			fmt.Printf("\r list[%d] is a different type", k)

		}
	}
}
