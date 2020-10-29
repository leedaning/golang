/*
* @Author: Leen
* @Date:   2020-03-09 16:44:37
* @Last Modified by:   Leen
* @Last Modified time: 2020-10-21 15:43:23
 */
package main

import (
	"fmt"
	"os"
)

func main() {
	/*s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}*/
	/*for i := 0; i < len(os.Args[0]); i++ {
		fmt.Println(i)
	}*/
	fmt.Println(os.Args[1])
}
