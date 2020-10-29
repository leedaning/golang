package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {

	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {

	rec := Rectangle{width: 2.12, height: 3.00}
	cir := Circle{2}

	fmt.Println("Rectangle area is", rec.area())
	fmt.Println("Circle area is", cir.area())
}
