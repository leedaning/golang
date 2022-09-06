package day7

import (
	"fmt"
	"math"
)

//自定义错误

func MyCustomize() {
	radius := -3.0
	area, err := CircleArea(radius)
	if err != nil {
		fmt.Println("错误信息：", err.Error())
	} else {
		fmt.Println("面积为：", area)
	}

}

//定义一个面积结构体，表示错误的类型
type areaError struct {
	msg    string
	radius float64 // 半径
}

//实现error接口，就是实现Error方法
func (r *areaError) Error() string {
	return fmt.Sprintf("error: 半径, %.2f, %s", r.radius, r.msg)
}

func CircleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"半径非法", radius} // areaError实现了Error()，故可以作为error类型的返回值
	}
	return math.Pi * radius * radius, nil
}
