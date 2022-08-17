package main

import (
	"fmt"
	"gostudy.com/day7"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello Leen")
	//second.MultipleFor()

	//MulitpleCircle()
	//DateFmt()
	//RandomNumInt()
	//MyArray()
	//mySlice()
	//mySlice2()
	//mySlice3()
	//day2.MyMap()
	//day2.MySlice()
	//day2.MyString()
	//day2.MyFunc()
	//day3.MyPoint()
	//day3.MyDefer()
	//day3.MyVarType()
	//day3.MyFuncType()
	//day3.MyAnonymity()
	//day3.MyCallBack()
	//day3.MyClosure()
	//day3.FuncVar()
	//day4.MyStruct()
	//day4.OOP()
	//day4.OOP2()
	//day4.Method()
	//day4.GetTree()
	//day5.MyInterface()
	//day5.MyItfe()
	//day5.MyNull()
	//day5.MyAssert()
	//day6.MyType()
	//day7.MyError()
	day7.MyErrorType()
}

func mySlice3() {
	var num int = 100
	fmt.Printf("num = %d, add = %p \n", num, &num)
	slc := make([]int, 2, 3)
	//slc := []int{}
	fmt.Println("\n切片slc:", slc)                                                                               // 切片slc: []
	fmt.Printf("切片长度:%d, 容量为：%d, 类型为：%T,  切片自身内存地址为：%p, 切片所指数组内存地址为：%p\n", len(slc), cap(slc), slc, &slc, slc) //切片长度:0, 容量为：0, 类型为：[]int,  内存地址为：0xe308d0

	slc = append(slc, 1)
	fmt.Println("\n切片slc:", slc)                                                     // 切片slc: [1]
	fmt.Printf("切片长度:%d, 容量为：%d, 类型为：%T,  内存地址为：%p\n", len(slc), cap(slc), slc, slc) //切片长度:1, 容量为：1, 类型为：[]int,  内存地址为：0xc0000044c0

	slc = append(slc, 2)
	fmt.Println("\n切片slc:", slc)                                                     // 切片slc: [1 2]
	fmt.Printf("切片长度:%d, 容量为：%d, 类型为：%T,  内存地址为：%p\n", len(slc), cap(slc), slc, slc) //切片长度:2, 容量为：2, 类型为：[]int,  内存地址为：0xc0000044c0

	slc = append(slc, 3)
	fmt.Println("\n切片slc:", slc)                                                     // 切片slc: [1 2 3]
	fmt.Printf("切片长度:%d, 容量为：%d, 类型为：%T,  内存地址为：%p\n", len(slc), cap(slc), slc, slc) //切片长度:3, 容量为：4, 类型为：[]int,  内存地址为：0xc0000044c0

	slc = append(slc, 4, 5, 6, 7, 8, 9)
	fmt.Println("\n切片slc:", slc)                                                     // 切片slc: [1 2 3 4 5 6 7 8 9]
	fmt.Printf("切片长度:%d, 容量为：%d, 类型为：%T,  内存地址为：%p\n", len(slc), cap(slc), slc, slc) //切片长度:9, 容量为：10, 类型为：[]int,  内存地址为：0xc000098440

}

func mySlice2() {
	slc := []int{}
	fmt.Println("\n切片slc:", slc)                                                                               // 切片slc: []
	fmt.Printf("切片长度:%d, 容量为：%d, 类型为：%T,  切片自身内存地址为：%p, 切片所指数组内存地址为：%p\n", len(slc), cap(slc), slc, &slc, slc) //切片长度:0, 容量为：0, 类型为：[]int,  切片自身内存地址为：0xc0000044c0, 切片所指数组内存地址为：0x1f08d0

	slc = append(slc, 1)
	fmt.Println("\n切片slc:", slc)                                                                               // 切片slc: [1]
	fmt.Printf("切片长度:%d, 容量为：%d, 类型为：%T,  切片自身内存地址为：%p, 切片所指数组内存地址为：%p\n", len(slc), cap(slc), slc, &slc, slc) //切片长度:1, 容量为：1, 类型为：[]int,  切片自身内存地址为：0xc0000044c0, 切片所指数组内存地址为：0xc0000140c0

	slc = append(slc, 2)
	fmt.Println("\n切片slc:", slc)                                                                               // 切片slc: [1 2]
	fmt.Printf("切片长度:%d, 容量为：%d, 类型为：%T,  切片自身内存地址为：%p, 切片所指数组内存地址为：%p\n", len(slc), cap(slc), slc, &slc, slc) //切片长度:2, 容量为：2, 类型为：[]int,  切片自身内存地址为：0xc0000044c0, 切片所指数组内存地址为：0xc0000140d0

	slc = append(slc, 3)
	fmt.Println("\n切片slc:", slc)                                                                               // 切片slc: [1 2 3]
	fmt.Printf("切片长度:%d, 容量为：%d, 类型为：%T,  切片自身内存地址为：%p, 切片所指数组内存地址为：%p\n", len(slc), cap(slc), slc, &slc, slc) //切片长度:3, 容量为：4, 类型为：[]int,  切片自身内存地址为：0xc0000044c0, 切片所指数组内存地址为：0xc00000a480

	slc = append(slc, 4, 5, 6, 7, 8, 9)
	fmt.Println("\n切片slc:", slc)                                                                               // 切片slc: [1 2 3 4 5 6 7 8 9]
	fmt.Printf("切片长度:%d, 容量为：%d, 类型为：%T,  切片自身内存地址为：%p, 切片所指数组内存地址为：%p\n", len(slc), cap(slc), slc, &slc, slc) //切片长度:9, 容量为：10, 类型为：[]int,  切片自身内存地址为：0xc0000044c0, 切片所指数组内存地址为：0xc0000121e0
}

func mySlice() {
	println("切片")
	//数组
	arr := [...]int{1, 2, 3}
	fmt.Println("数组arr为：", arr)

	//切片
	var slc []int
	fmt.Println("切片slc为：", slc)

	//变长的数组切片
	slc2 := []int{1, 2, 3}
	fmt.Println("切片slc2为：", slc2)
	fmt.Printf("切片slc2类型为：%T\n", slc2)
	fmt.Printf("数组arr类型为：%T\n", arr)

	slc3 := make([]int, 3, 6)
	fmt.Println("切片slc3为：", slc3)
	fmt.Printf("切片slc3类型为：%T\n", slc3)
	fmt.Printf("切片slc3, 长度为：%d, 容量为：%d\n", len(slc3), cap(slc3))

	slc4 := make([]int, 0, 5)
	fmt.Println("\n切片slc4为：", slc4)
	fmt.Printf("切片slc4类型为：%T\n", slc4)
	slc4 = append(slc4, 1, 2) // 给切片增加元素，要用append，超出切片长度要重新创建一个切片，并赋给原切片
	fmt.Println("切片slc4为：", slc4)
	fmt.Printf("切片slc4类型为：%T\n", slc4)
	slc4 = append(slc4, slc2...) // 将切片slc2增加到切片slc4中
	fmt.Println("切片slc4为：", slc4)
	fmt.Printf("切片slc4类型为：%T\n", slc4)

	for key, value := range slc4 {
		//fmt.Printf("\n key:", key, " value:", value)
		fmt.Println("\n key:", key, " value:", value)
	}

}

func MyArray() {
	/*var months [6]int
	months[0] = 1
	months[1] = 2
	months[2] = 3
	months[3] = 4
	months[4] = 5
	months[5] = 6*/
	//var months = [6]int{1}
	//months := [...]int{0: 1, 1: 2}
	//var months = [...]int{1, 2, 3, 4, 5, 6}
	months := [2]struct {
		key   int
		value int
	}{
		{0, 1},
		{1, 2},
	}

	fmt.Println("月份：", months)
	for key, value := range months {
		fmt.Printf("\n key=%d,  value=%d;", key, value)
	}
}

//随机数
func RandomNumInt() int {
	//num := rand.Int()
	rand.Seed(22)
	num := rand.Intn(90)
	fmt.Println("生成的随机数为：", num)
	return num
}

func DateFmt() {
	timeLayout := "2006-01-02 15:04:05"

	stime, err := time.Parse(timeLayout, "2022-06-30 12:10:00")
	if err == nil {
		fmt.Printf("\n%T\n", stime)
		fmt.Println("stime：", stime)

		unixTime := stime.Unix()
		fmt.Printf("\n%T\n", unixTime)
		fmt.Println("unixTime：", unixTime)

		timeTime := time.Unix(unixTime, 0)
		fmt.Printf("\n%T\n", timeTime)
		fmt.Println("timeTime：", timeTime)

		dateTime := timeTime.Format("2006-01-02")
		fmt.Printf("\n%T\n", dateTime)
		fmt.Println("日期：", dateTime)

		dateTime2 := stime.Format("2006-01-02")
		fmt.Printf("\ndateTime2 类型：%T\n", dateTime2)
		fmt.Println("dateTime2：", dateTime2)

		stime2, _ := time.Parse("2006-01-02", dateTime2)
		fmt.Printf("\nstime2 类型：%T\n", stime2)
		fmt.Println("stime2：", stime2)

		//unixTime2 := time.Unix(dateTime2, 0)
	}
}

func MulitpleCircle() {

out:
	for i := 1; i <= 10; i++ {
		fmt.Printf("i=%d \t", i)
		for j := 1; j <= 10; j++ {

			if i == 2 {
				fmt.Println()
				continue out
			}
			fmt.Printf("%d \t", j)
		}
		fmt.Println()
	}

	i := 1
	for ; i <= 10; i++ {
		if i == 6 {
			i++
			goto breakHere
		}
		fmt.Println("i = ", i)

	}

breakHere:
	fmt.Println("i的最终结果为：", i)
}
