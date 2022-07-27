package day3

import "fmt"

func MyPoint() {

	fmt.Println("指针")
	DefinePoint()
	//ArrPoint()
	//pointArr()
	/*arr1 := createArr()
	fmt.Printf("arr1的内存地址为：%p，本身的值为：%p\n", &arr1, arr1)
	arr2 := createArr()
	fmt.Printf("arr2的内存地址为：%p，本身的值为：%p\n", &arr2, arr2)*/
	//pp()
}

//指针的指针
func pp() {
	var num int = 100         // int类型的变量
	var arr = [3]int{1, 2, 3} // int类型的数组
	var slc = []int{}         // int类型的切片
	var pArr = [1]*int{}      // 指针数组

	var p1 *int = &num    // 指针
	var p2 **int = &p1    // 指针类型的指针
	var p3 *[3]int = &arr // 数组指针
	var p4 *[]int = &slc  // 切片指针
	pArr[0] = p1
	var funP func() *int // 函数指针
	funP = funPoint      // 函数类型
	p5 := funP()         // 指针类型

	fmt.Println("变量：")
	fmt.Printf("num = %d, type = %T, add = %p \n", num, num, &num)
	fmt.Printf("arr = %d, type = %T, add = %p \n", arr, arr, &arr)
	fmt.Printf("slc = %d, type = %T, add = %p \n", slc, slc, &slc)
	fmt.Printf("pArr = %p, type = %T, add = %p \n", pArr, pArr, &pArr)
	fmt.Println("pArr：", pArr)

	fmt.Println("\n指针：")
	fmt.Printf("p1 = %p, type = %T, add = %p \n", p1, p1, &p1)
	fmt.Printf("p2 = %p, type = %T, add = %p \n", p2, p2, &p2)
	fmt.Printf("p3 = %p, type = %T, add = %p \n", p3, p3, &p3)
	fmt.Printf("p4 = %p, type = %T, add = %p \n", p4, p4, &p4)
	fmt.Printf("funP = %p, type = %T, add = %p \n", funP, funP, &funP) // 函数类型
	fmt.Printf("p5 = %p, type = %T, add = %p \n", p5, p5, &p5)         // 函数指针，指针类型
}

//指针函数，返回值是指针的函数
func funPoint() *int {
	var number int = 100
	fmt.Printf("number = %d, type = %T, add = %p \n", number, number, &number)

	return &number
}

//创建数组
func createArr() [3]int {
	arr := [3]int{11, 22, 33}
	fmt.Println("arr:", arr)
	fmt.Printf("arr的内存地址为：%p\n", &arr)
	return arr
}

//指针数组
func pointArr() {
	var arr [3]*int
	var a, b, c int = 10, 20, 30
	fmt.Printf("数组arr的内存地址为：%p, 类型为：%T\n", &arr, arr)
	fmt.Println("a = %d, b = %d, c = %d\n", a, b, c)
	fmt.Printf("a的地址：%p, b的地址：%p, c的地址：%p \n", &a, &b, &c)
	fmt.Println("数组arr的值为：", arr)
	arr[0] = &a
	arr[1] = &b
	arr[2] = &c
	fmt.Println("数组arr的值为：", arr)
	a = 100
	fmt.Println("数组arr的值为：", arr)
	for k, v := range arr {
		fmt.Println("k = ", k, " v = ", *v)
	}
}

//数组指针
func ArrPoint() {

	//var arrPoint *[3]int
	arr := [3]int{1, 2, 3}
	arrPoint := &arr

	fmt.Printf("数组指针的类型为：%T，内存地址为：%p\n", arrPoint, &arrPoint)
	fmt.Println("数组指针arrPoint:", arrPoint)

}

//指针声明、使用
func DefinePoint() {

	var p *int
	var num int = 100
	p = &num
	var p2 *int
	var num2 int = 200
	p2 = &num2
	fmt.Println("num=", num, " 地址为：", &num)
	//fmt.Printf("变量num的值为：%d， 地址为：%p\n", num, &num)
	fmt.Printf("指针本身的内存地址为:%p，指针p的值为：%p, 指针值所指的地址的值为：%d", &p, p, *p)
	ChangeVal(p)
	fmt.Printf("\n\n变量num的值为：%d， 地址为：%p\n", num, &num)
	fmt.Printf("指针本身的内存地址为:%p，指针p的值为：%p, 指针值所指的地址的值为：%d", &p, p, *p)
	ChangeVal(p)
	fmt.Printf("\n变量num的值为：%d， 地址为：%p\n", num, &num)
	fmt.Printf("指针本身的内存地址为:%p，指针p的值为：%p, 指针值所指的地址的值为：%d", &p, p, *p)

	fmt.Println("\n\nnum=", num2, " 地址为：", &num2)
	//fmt.Printf("\n变量num2的值为：%d， 地址为：%p\n", num2, &num2)
	fmt.Printf("指针本身的内存地址为:%p，指针p的值为：%p, 指针值所指的地址的值为：%d", &p2, p2, *p2)
	ChangeVal(p2)
	fmt.Printf("\n变量num2的值为：%d， 地址为：%p\n", num2, &num2)
	fmt.Printf("指针本身的内存地址为:%p，指针p的值为：%p, 指针值所指的地址的值为：%d", &p2, p2, *p2)

}

func ChangeVal(num *int) {
	//*num = *num + 1
	*num++
}
