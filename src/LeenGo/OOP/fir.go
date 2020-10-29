package main

import (
	"fmt"
)

type Rectangle struct {
	width, height float32
}

func area(r Rectangle) float32 {
	return r.width * r.height
}

func main() {
	fir := Rectangle{width: 5.0, height: 10.00}
	sec := Rectangle{width: 6.50, height: 20.00}

	fmt.Println("fir area:", area(fir))
	fmt.Println("sec area:", area(sec))
}
