package main

import (
	. "fmt"
	. "math"
)

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, hight, deepth float64
	color                Color
}

type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.hight * b.deepth
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)

	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack() {
	for i, v := range bl {
		Printf("\r\r bl i is : %v, v is %v", i, v)
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strs := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strs[c]
}

func main() {

	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	Println("The boxes info :", boxes)
	Println("The π value is ", Pi)
	Println("The biggest color is :", boxes.BiggestColor())
	Println("The biggest color is :", boxes.BiggestColor().String())

	boxes.PaintItBlack()

	Println("The boxes info :", boxes)
	Println("The biggest color is :", boxes.BiggestColor())
}
