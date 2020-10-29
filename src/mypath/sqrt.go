/*
* @Author: Leen
* @Date:   2020-01-07 20:11:59
* @Last Modified by:   Leen
* @Last Modified time: 2020-01-07 20:12:10
 */
// package mymath
package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 0.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}
func main() {
	fmt.Println("Leen")
	fmt.Println(Sqrt(100))
}
