package main

import (
	"fmt"
)

var intf interface{}

var Num int = 5
var Email string = "leening@163.com"

func main() {
	Name := "leen"
	fmt.Printf("Num = %d, Name = %s", Num, Name)
	Num = add()
	fmt.Printf("\rNum = %d, Name = %s", Num, Name)

	made()
}

func add() int {
	Num++
	return Num
}

func made() {
	intf = Num
	fmt.Printf("\r intf = %v", intf)

	intf = Email
	fmt.Printf("\r intf = %v", intf)

}
