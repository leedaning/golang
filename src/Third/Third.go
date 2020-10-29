/*
* @Author: Leen
* @Date:   2020-03-09 17:14:57
* @Last Modified by:   Leen
* @Last Modified time: 2020-03-09 17:21:42
*/
package main

import (
	"fmt"
	"os"
)

func main(){
    var s, sep string
    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    fmt.Println(s)
}