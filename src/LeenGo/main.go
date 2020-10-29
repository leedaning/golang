// LeenGo project main.go
package main

import (
	"errors"
	"fmt"
)

func main() {
	// export()
	// str()
	// err()
	// group()
	// Arr()
	// Slice()
	// myMap()

	/*num := 1
	step := 2
	increase(num, step)*/

	// conditions()
	mySwitch()
}

func export() {
	var a int = 10
	var b, c, d int = 1, 2, 3
	e, f := 4, 5

	fmt.Println("Hello World!")
	fmt.Print('A')
	fmt.Printf("a=%v", a)
	fmt.Sprintf("b=%d", b)
	fmt.Sprintf("c=%d", c)
	fmt.Sprintf("d=%d", d)
	fmt.Sprintf("e=%d", e)
	fmt.Sprintf("e=%d", e)
	fmt.Sprintf("f=%d", f)
}

func str() {
	str := `\r\nThis is
                    string type.`
	fmt.Printf(str)
}

func err() {
	err := errors.New("\remit macho dwarf: elf header corrupted\n")
	if err != nil {
		fmt.Print(err)
	}
}

func group() {
	const (
		Pi   = 3.14159226
		Leen = 100
	)
	fmt.Printf("\rPi=%f", Pi)
	fmt.Printf("\rLeen=%d", Leen)
}

func Arr() {
	var arr [10]int
	fruits := [6]string{"banara", "apple", "orange", "Pear", "Sugar cane", "Grape"}
	peoples := [2][3]string{[3]string{"Joy", "LiLei", "John"}, [3]string{"Lucy", "Jan", "Nv"}}

	arr[0] = 0
	arr[1] = 100

	fmt.Printf("\r arr 0=%d", arr[0])
	fmt.Printf("\r arr 1=%d", arr[1])
	fmt.Printf("\r Have so many fruit, eg. %v、 %v、 %v、 %v、 %v、 %v", fruits[0], fruits[1], fruits[2], fruits[3], fruits[4], fruits[5])
	fmt.Printf("\r male: %v", peoples[0])
	fmt.Printf("\r female: %v", peoples[1])
	fmt.Printf("\r The array fruites length is %d", len(fruits))
}

func Slice() {
	alphabet := [27]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var fslice, firArr, secArr []string // 声明一个slice，跟声明数组类似，只是没有长度
	slice := []byte{'a', 'b', 'c', 'd'} // 声明一个slice，并初始化数据

	fslice = alphabet[0:26]
	firArr = alphabet[1:3]
	secArr = alphabet[3:4]

	fmt.Printf("\rThis is slice : %v", slice)
	fmt.Printf("\rThis is fslice : %v", fslice)
	fmt.Printf("\rThis is firArr : %v", firArr)
	fmt.Printf("\rThis is secArr : %v", secArr)
	fmt.Printf("\rThe alphabet have %d number!", len(alphabet))
	fmt.Printf("\rThe alphabet have %d number!", cap(alphabet))

	slice = append(slice, 'A')
	fslice = append(fslice, "A")

	fmt.Printf("\rThis is slice : %v", slice)
	fmt.Printf("\rThis is fslice : %v", fslice)

}

func myMap() {
	// map的第一种声明方式
	var myMap map[int]string
	myMap = make(map[int]string)
	myMap[0] = "Great"
	myMap[1] = "Beautify"
	myMap[2] = "Optimistic"

	// var otherMap map[int]string
	// otherMap = make(map[int]string)
	otherMap := myMap
	otherMap[0] = myMap[0]
	otherMap[1] = "Good"

	fmt.Printf("\rThis is a %s day! len:%d", myMap[0], len(myMap))
	fmt.Printf("\rThis otherMap one is :%v, second:%v ", otherMap[0], otherMap[1])
	fmt.Printf("\r This is otherMap:%v", otherMap)

	fmt.Printf("\rmyMap:%v", myMap)
	otherMap[0] = "Very good!"
	fmt.Printf("\r This is myMap:%v", myMap)
	fmt.Printf("\r This is otherMap:%v", otherMap)

	// map的第二种声明方式
	secMap := map[int]string{0: "lee", 1: "daning"}
	delete(secMap, 0) //输出key为0的元素
	name, ok := secMap[0]
	fmt.Printf("\r\r secMap:%v", secMap)
	fmt.Printf("\r name:%v, ok:%v", name, ok)

}

// func increase(num, step) {
// 	return num + step
// }

/*
流程控制
*/
func conditions() {

	var total int = 1

	// if条件判断，一
	if total > 5 {
		fmt.Printf("\rTotal number > 5, is %d", total)
	} else {
		fmt.Printf("\rTotal num < 5")
	}

	// if条件判断，二
	if num := 10; num > 5 {
		fmt.Printf("\rnum > 5")
	} else {
		fmt.Printf("\rnum < 5")
	}

	// Goto
Here:

	fmt.Printf("\rtotal:%d", total)
	total++
	if total <= 10 {
		goto Here // 此处一定要小心，不然会无限循环下去的
	}

	fmt.Print("\r\rFor:\r")
	// for
	var totalNum = 0
	for index := 1; index <= 10; index++ {
		totalNum += index
		fmt.Printf("\rindex:%d", index)
	}
	fmt.Printf("\rtotalNum:%d", totalNum)

	// while功能
	fmt.Print("\r\rWhile:\r")
	sum := 1
leen:
	for sum <= 10 {
		fmt.Printf("\rSum:%d", sum)
		for total := 1; total <= 10; total++ {

			if sum > 5 {
				break leen
			}
			fmt.Printf("\ttotal:%d", total)
		}
		sum++
	}

	// for 用于读取slice和map的数据
	var maps map[int]string
	maps = make(map[int]string)
	maps[0] = "Today"
	maps[1] = "is"
	maps[2] = "good"
	maps[3] = "day"

	for k, v := range maps {
		fmt.Println("\rmap's key:", k)
		fmt.Println("\rmap's val:", v)
	}

	for _, v := range maps {
		fmt.Println("\r\r Map's val:", v)
	}
}

func mySwitch() {
	var fruit string
	fruit = "fruit"

	switch fruit {
	case "apple": //go的case后面自带break
		fmt.Print("\rThis is apple")
		fallthrough // 强制不跳出switch，继续执行
	case "banana", "fruit": // 匹配多个值
		fmt.Print("\rThis is banana")
	default:
		fmt.Print("This is fruit")
	}
}
