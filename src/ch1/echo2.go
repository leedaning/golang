/*
* @Author: Leen
* @Date:   2019-08-27 20:21:54
* @Last Modified by:   Leen
* @Last Modified time: 2019-08-27 20:26:29
*/
package main

import(
	"fmt"
	"os"
)

func main(){
	s, sep := "", ""
	for _, arg := range os.Args[1:]{
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
