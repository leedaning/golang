package day4

import "fmt"

func GetTree() {

	/*var list []map[string]int
	list = append(list, map[string]int{"id": 1, "pid": 0})
	list = append(list, map[string]int{"id": 2, "pid": 0})*/

	/*list := make([]int, 0)
	list = append(list, 10, 11)
	list = append(list, 20, 21)*/

	list := make([]map[string]int, 0)
	list = append(list, map[string]int{"id": 1, "pid": 0})
	list = append(list, map[string]int{"id": 2, "pid": 0})
	list = append(list, map[string]int{"id": 3, "pid": 0})
	list = append(list, map[string]int{"id": 4, "pid": 0})
	list = append(list, map[string]int{"id": 5, "pid": 0})
	list = append(list, map[string]int{"id": 6, "pid": 0})
	list = append(list, map[string]int{"id": 7, "pid": 0})
	list = append(list, map[string]int{"id": 8, "pid": 0})
	list = append(list, map[string]int{"id": 9, "pid": 0})
	list = append(list, map[string]int{"id": 10, "pid": 0})

	fmt.Println("list:", list)
}
