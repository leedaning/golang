package day3

import "fmt"

func MyDefer() {
	//defer延迟执行，先加载函数或者方法，不执行，等到程序全部执行结束的时候再执行

	var num int = 10

	fmt.Printf("\nnum 的原值为：%d", num)
	defer Add(num)
	num++
	//defer Add(100)
	fmt.Printf("\nnum 的值为：%d", num)

}
func Add(num int) int {
	//num++
	fmt.Printf("\nAdd()中的值为：%d", num)
	return num
}
