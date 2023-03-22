package day8

import (
	"fmt"
	"time"
)

func Concurrency() {
	fmt.Println(10 * time.Millisecond)
	fmt.Println(time.Duration(1) * time.Second)
	/*for i := 0; i <= 10; i++ {
		go MyPrt(i)
		fmt.Println(" i i = ", i)
	}
	time.Sleep(time.Duration(1) * time.Second)*/

}

func MyPrt(i int) {
	fmt.Println("i = ", i)
}
