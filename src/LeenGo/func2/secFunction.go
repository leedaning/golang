package main

import "fmt"

/*
type testInt func(int) bool // 声明了一个函数类型

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

// 声明的函数类型在这个地方当做了一个参数

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd) // 函数当做值来传递了
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven) // 函数当做值来传递了
	fmt.Println("Even elements of slice are: ", even)
}
*/

type myFunc func(int) bool // 声明一个函数类型
func isOdd(num int) bool {
	if num%2 == 0 {
		return false
	}
	return true
}

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f myFunc) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("\r My slice is ", mySlice)
	odd := filter(mySlice, isOdd)
	fmt.Println("\rThe odd is :", odd)
	even := filter(mySlice, isEven)
	fmt.Println("\rThe even is :", even)
}
