package main

import (
	"fmt"
)

type intf1 interface {
	run()
}

type intf2 interface {
	intf1
	sing()
}

type Human struct {
	name string
	age  int
}

func (h Human) run() {
	fmt.Printf("\rThe human run , name:%s, age:%d", h.name, h.age)
}

func (h Human) sing() {
	fmt.Printf("\rThe human sing")
}

func main() {
	human := Human{"leen", 18}

	human.run()
	human.sing()

	var intfOne intf1
	var intfTwo intf2

	intfOne = human
	intfOne.run()

	intfTwo = human
	intfTwo.run()
	intfTwo.sing()

}
