package main

import (
	"fmt"
	"os"
)

var user = os.Getenv("USER")

func init() {

	fmt.Println("\rThis is init function.")
	// if user == "" {
	// 	panic("no value for $USER")
	// }
}

func main() {

	fmt.Println("Index access main function.")
	throwsPanic(firPainc)
}

func getEnv() {

	var JAVA_HOME = os.Getenv("JAVA_HOME")

	fmt.Println("JAVA_HOME is :", JAVA_HOME)
}

func firPainc() {
	fmt.Println("\rThis is firPainc function.")
	panic("\rHere is a painc.")
}

func throwsPanic(f func()) (b bool) {
	defer func() {
		fmt.Println("\rb的值为", b)
		if x := recover(); x != nil {
			fmt.Println("\rx的值为", x)
			b = true
		}
	}()
	f() //执行函数f，如果f中出现了panic，那么就可以恢复回来
	return
}
