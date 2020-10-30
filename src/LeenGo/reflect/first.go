package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "leedaning"

	v := reflect.ValueOf(name)

	fmt.Printf("\rType is %v", v.Type())                         // 变量类型
	fmt.Printf("\rValue is %s", v.String())                      // 变量值
	fmt.Println("\rkind is string:", v.Kind() == reflect.String) // 比较类型数值跟当前类型数值是否相等

	// 修改值
	p := reflect.ValueOf(&name)
	s := p.Elem()
	s.SetString("leen")
	fmt.Printf("\r\r Value is %s", s.String())
}
