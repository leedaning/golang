package day5

import "fmt"

//接口断言

//植物接口
type Plant interface {
	life() // 有生命
}

//花类接口
type Flower interface {
	Plant
	flowering() // 会开花
}

//草类接口
type Grass interface {
	Plant
	vitality() // 生命力，顽强
}

//花类的结构体
type Flowers struct {
	name string
}

//草类的结构体
type Grasses struct {
	name string
}

//花类结构体实现life()方法，实现Plant接口
func (r Flowers) life() {
	fmt.Println(r.name + "是有生命的")
}

//花类结构体实现flowering()方法，实现Flower接口
func (r Flowers) flowering() {
	fmt.Println(r.name + "会开花")
}

//草类结构体实现life()方法，实现Plant接口
func (r Grasses) life() {
	fmt.Println(r.name + "是有生命的")
}

//草类结构体实现vitality()方法，实现Flower接口
func (r Grasses) vitality() {
	fmt.Println(r.name + "生命力顽强")
}

func MyAssert() {
	var flowers = Flowers{name: "迎春花"}
	/*flowers.life()
	flowers.flowering()*/

	var grasses = Grasses{"狗尾草"}

	Assert(flowers)

	//声明一个接口类型
	var plant Plant
	plant = flowers
	fmt.Printf("花类接口变量内存地址：%p\n", &plant)
	Assert(plant)

	var plant2 Plant
	plant2 = grasses
	fmt.Printf("草类接口变量内存地址：%p\n", &plant2)
	Assert(plant2)

	//声明一个指针类型的接口变量
	var flowerPoint *Flowers = &Flowers{name: "迎春花"}
	fmt.Printf("花类结构体指针变量内存地址：%p\n", flowerPoint)
	Assert(flowerPoint)

}

func Assert(p Plant) {

	fmt.Println(p)
	fmt.Printf("变量内存地址：%p\n", &p)

	//第一种断言方式，遇到断言失败会造成一个painc，如果p为nil同样也会panic，如：panic: interface conversion: day5.Grasses is not day5.Flower: missing method flowering
	/*ins := p.(Flower)
	fmt.Println("断言结果：", ins)*/

	//第二种断言方式，断言失败不会造成panic
	if ins, ok := p.(Flower); ok {
		fmt.Println("花类 ins:", ins, "  --  ok:", ok)
	} else if ins, ok := p.(*Flowers); ok { // 接口指针类型的断言
		fmt.Println("花类指针 ins:", ins, "  --  ok:", ok, "-- 内存地址：%p", &ins)
	} else if ins, ok := p.(Grass); ok {
		fmt.Println("ins:", ins, "  --  ok:", ok)
	} else {
		fmt.Println("断言失败")
	}

	//第三种断言方式
	/*switch p.(type) {
	case Flower:
		fmt.Println("花类接口断言成功")
		p.life()
	case Flowers:
		fmt.Println("花类结构体断言成功")
		p.life()
	case *Flowers: // 断言这里的地址跟调用的接口变量的内存地址是不一样的
		fmt.Println("花类结构体指针断言成功")
		//fmt.Printf("变量内存地址：%p\n", &p)
		p.life()
	case Grass:
		fmt.Println("草类断言成功")
		p.life()
	default:
		fmt.Println("断言失败")
	}*/

	fmt.Println("||||||")

}
