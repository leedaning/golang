package main

import (
	"fmt"
)

// var a int = 100
// var b int = 200
// var t int
var a, b byte = 'a', 'b'
var t byte

func main() {
	// exchange()
	exchange2()
}

func exchange() {
	t = a
	a = b
	b = t
	fmt.Println(a, b)
}

func exchange2() {
	fmt.Println(a, b)

	a = a ^ b
	b = a ^ b
	a = b ^ a

	fmt.Println(a, b)
}
