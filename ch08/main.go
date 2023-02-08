package main

import "fmt"

type Person struct {
	name string
}

// 接受者
func (p *Person) SayHello() {

}

// 通过指针交换两个值
func swap(a, b *int) {
	// a篮子中存放b篮子中的值，b篮子中存放a篮子的值
	// 零时值
	t := *a

	// 将a地址的值，改为b指向的值
	*a = *b
	*b = t
}

// 指针类型接受的是地址
func changeName(p *Person) {
	p.name = "张三"
}

func main() {

	// 指针，提一个需求，我希望结构体传值的时候，我在函数中修改的值可以返回到变量中去
	// 这样定义p是值类型不是指针类型，如果定义指针类型可以用
	// p := &Person{name: "xmaven"}
	p := Person{name: "xmaven"}
	// p 是对象，&是取地址
	changeName(&p)
	fmt.Println(p)

	var pi *Person = &p
	// fmt 打印指针对指针进行了优化处理，如果想打印地址
	fmt.Println(pi)
	fmt.Printf("%p\r\n", pi)

	// 指针的定义
	var po *Person
	po = &p
	// 如何修改指针对象的值 第一种
	(*po).name = "王武"
	fmt.Println(po)
	// 如何修改指针对象的值，第二种 go独有的，c和c++没有的
	po.name = "李四"
	// 第一个不同的点就出来了，第二点 go语言限制了指针的运算，在C语言中你拿到了一个指针可以进行+1，在go语言中不行，不能参加运算
	//	go指针就是一个阉割版，unsafe包里面，所以说一般我们不会使用unsafe包，当你要使用的时候是可以使用的
	fmt.Println(po)

	// 指针的初始化
	var a int = 10
	b := &a
	fmt.Println(b)

	fmt.Println("11111111111111111111111111111111")

	// 指针需要初始化
	var c *int
	fmt.Println(c)

	//结构体指针初始化
	// 空指针无法访问属性 必须需要初始化
	//var p2 *Person
	//fmt.Println(p2.name)

	fmt.Println("2222222222222222222222222222222222")
	// 第一种初始化方式
	p3 := &Person{}
	fmt.Println(p3.name)

	fmt.Println("33333333333333333333333333333333333")
	// 第二种初始化方式
	var emptyPerson Person //定义的时候就已经初始化好了
	p4 := &emptyPerson
	fmt.Println(p4.name)

	fmt.Println("44444444444444444444444444444444444444444444")

	// 初始化两个关键字，map channel slice 初始化推荐使用make方法
	// 指针初始化推荐使用new函数，指针要初始化否则会出现nil point
	// map 必须初始化

	// 第三种初始化方式
	var pp *Person = new(Person)
	fmt.Println(pp.name)

	// 交换两个值
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)

}
