package main

import (
    "fmt"
)

type myFunc func(int) bool //自定义函数类型

/*
 是否是奇数
*/
func isOdd(num int) bool {
    if num%2 == 0 {
        return false
    }
    return true
}

/*
 是否是偶数
*/
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
    fmt.Println("\rSlice :", mySlice)
    odd := filter(mySlice, isOdd)
    fmt.Println("\rOdd :", odd)
    even := filter(mySlice, isEven)
    fmt.Println("\rEven :", even)
}
