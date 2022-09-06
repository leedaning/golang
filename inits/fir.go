package inits

import "fmt"

func init() {
	fmt.Println("This is fir file init() one.")
}

func Start() {
	fmt.Println("init相关， start")
}

func init() {
	fmt.Println("This is fir fiel init() two.")
}
