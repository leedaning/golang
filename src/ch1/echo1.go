/*
* @Author: Leen
* @Date:   2019-08-27 20:00:30
* @Last Modified by:   Leen
* @Last Modified time: 2020-03-09 16:44:19
*/
package main

import (
    "fmt"
    "os"
)

func main() {
    var s, sep string
    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    fmt.Println(s)
}