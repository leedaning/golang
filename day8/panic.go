package day8

import (
	"fmt"
	"log"
)

//panic recover
func MyPanic() {

	//go func() {
	defer func() {
		fmt.Println("入口")
		if err := recover(); err != nil {
			log.Printf("recover:%v", err)
		}
	}()
	//}()

	defer Prt("01")
	CirclePrt()
	//panic("中断") // 中断程序的执行，不再继续往下执行，前面调用的defer语句还是会执行，后面的defer语句则不会执行
	defer Prt("02")
}

func Prt(name string) {
	fmt.Println("这是打印函数，打印：", name)
}

func CirclePrt() {
	defer Prt("03")
	panic("中断CirclePrt方法") // 中断程序的执行，不再继续往下执行，前面调用的defer语句还是会执行，后面的defer语句则不会执行
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}

	defer Prt("04")
}
