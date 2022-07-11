package day2

import (
	"fmt"
	"sort"
)

func MyMap() {
	fmt.Println("Hello map")
	//MapDefine()
	//MapOprt()
}

func MapOprt() {

	var myMap = make(map[int]string)

	if myMap == nil {
		fmt.Println("myMap 为空")
		myMap = make(map[int]string)
	}

	myMap[1] = "Leen"
	myMap[2] = "John"
	myMap[3] = "Lucy"
	myMap[4] = "Tom"
	myMap[5] = "Jim"
	fmt.Println(myMap)

	//输出是无序的
	for k, v := range myMap {
		println(k, ":", v)
	}

	//有序输出，需要借助切片
	slc := make([]int, 0)
	for k, _ := range myMap {
		slc = append(slc, k)
	}
	fmt.Println("切片为：", slc)
	sort.Ints(slc)
	fmt.Println("切片排序后为：", slc)
	fmt.Println("map有序输出")
	for _, v := range slc {
		println("index=", v, "\t value=", myMap[v])
	}

	/*val, ok := myMap[20]
	if ok {
		println("val :", val)
	} else {
		println("操作的key不存在，获取到的是零值：", val)
	}*/

	/*var myMap2 = make(map[string]string)
	myMap2["name"] = "leen"
	myMap2["age"] = "young"
	myMap2["sex"] = "male"
	myMap2["heigh"] = "170"
	fmt.Println(myMap2)
	for k, v := range myMap2 {
		println(k, ":", v)
	}*/

	/*var myMap2 map[int]string
	fmt.Printf("\nmyMap2的类型为:%T, 地址为：%p, 引用地址为：%p\n", myMap2, &myMap2, myMap2)
	fmt.Println("myMap2的内容为：", myMap2)

	if myMap2 == nil {
		fmt.Println("myMap2 为空")
	}
	fmt.Println(myMap2 == nil)
	if myMap2 == nil {
		myMap2 = make(map[int]string)
		fmt.Printf("\nmyMap2的类型为:%T, 地址为：%p, 引用地址为：%p\n", myMap2, &myMap2, myMap2)
		fmt.Println("myMap2的内容为：", myMap2)
	}*/

}

//map的声明
func MapDefine() {
	var map1 map[int]string // 这种声明只有声明，没有初始化，不可以用map[key]这种形式赋值，会报错

	fmt.Printf("\nmap1的类型为:%T, 地址为：%p, 引用地址为：%p\n", map1, &map1, map1)
	fmt.Println("map1的内容为：", map1)

	var map2 = make(map[int]string, 5)
	fmt.Printf("\nmap2的类型为:%T, 地址为：%p, 引用地址为：%p\n", map2, &map2, map2)
	map2[0] = "Hello world"
	map2[1] = "Hello everyone"
	map2[2] = "Hello Leen"
	fmt.Println("map2的内容为：", map2)

	var map3 = map[int]string{0: "name", 1: "sex", 2: "mobile"}
	fmt.Printf("\nmap3的类型为:%T, 地址为：%p, 引用地址为：%p\n", map3, &map3, map3)
	fmt.Println("map3的内容为：", map3)

	map3[1] = "age"
	fmt.Printf("\nmap3的类型为:%T, 地址为：%p, 引用地址为：%p\n", map3, &map3, map3)
	fmt.Println("map3的内容为：", map3)

	map4 := map[int]string{0: "name", 1: "sex", 2: "mobile"}
	fmt.Printf("\nmap4的类型为:%T, 地址为：%p, 引用地址为：%p\n", map4, &map4, map4)
	fmt.Println("map4的内容为：", map4)
}
