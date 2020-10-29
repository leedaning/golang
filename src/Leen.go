package main

import "fmt"

func main() {
	fmt.Println("Hello Leen!")
	var a int
	// Scan()等待键盘输入。  &表示取地址
	fmt.Scan(&a)   // 如果输入的不是int类型，那么a就是默认的0
	fmt.Println(a) // 0
}



