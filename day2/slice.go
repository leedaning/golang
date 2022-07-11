package day2

import "fmt"

func MySlice() {
	mySlice()
}

func mySlice() {
	mySlc := []string{"name", "age", "sex", "", "heigh", "weight"}
	fmt.Printf("mySlc Type : %T \n", mySlc)
	fmt.Println("mySlc:", mySlc)
	for k, v := range mySlc {
		fmt.Println(k, ":", v)
	}

	myMap := make(map[int]string)
	fmt.Printf("\nmyMap Type : %T \n", myMap)
	fmt.Println("myMap:", myMap)

	for k, v := range mySlc {
		myMap[k] = v
	}
	fmt.Println("myMap:", myMap)

	val, ok := myMap[10]
	if ok {
		fmt.Println("map对应的下标的存在", val)
	} else {
		fmt.Println("map对应的下标的值不存在")
	}
	//fmt.Println("mySlc[1]", mySlc[20])
	/*val, ok := mySlc[20]
	if ok {

	}*/
}
