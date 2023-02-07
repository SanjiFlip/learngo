package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 结构体嵌套
type Person struct {
	name string
	age  int
}

type Student struct {
	// 第一种嵌套方式
	p     Person
	score float32
}

type Student2 struct {
	// 第二种嵌套方式 匿名嵌套
	Person
	score float32
}

type Human struct {
	//java和c++中的可访问性 public、 private、protected
	//python 没有的， 可访问，不可访问 规范， 可访问性 针对的是struct的字段、函数、变量等
	Name   string `json:"name,omitempty"`        //一些情况下很有用 omitempty表示忽略空值
	Age    int    `json:"age,string" bobby:"go"` //很好用， 其他语言中 java， python， php，实际上你想要做到这个很简单
	Height int    `json:"-"`                     // 表示没有赋值或者没有默认值就不展示
	//那鬼知道这些有哪些写法呢 本质就是一个tag的字符串
}

// 结构体定义方法
// func (s StructType)funcName(param1 param1Type, ...)(returnVal1 returnType1, ...){ }
// 接收器有两种形态
func (p Person) print() {
	fmt.Printf("name:%s, age:%d\r\n", p.name, p.age)
}

func (p *Person) print2() {
	// 有可能该函数中想要改变p的值，就是person对象很大，数据很大的时候可以考虑指针传递
	p.age = 110
	fmt.Printf("name:%s, age:%d\r\n", p.name, p.age)
}

// 空结构体不占空间
type MySet map[string]struct{}

//go语言没有面向对象一些class的概念，但是代码的封装，面向对象的思想还有的，
//面向对象的3大特性： 封装、继承 和多态(接口)
//把字段和方法封装到一起 这个方法可以操作这个class上的字段

func main() {
	s := Student{
		Person{
			"xmaven", 19,
		},
		95.7,
	}
	s2 := Student{}
	s2.p.name = "bystart"
	fmt.Println(s.p.age)
	fmt.Println(s.p.name)

	fmt.Println("----------------")
	// 匿名嵌套可以直接访问
	/*
		赋值和取值的时候都可以直接访问
		但是初始化的时候不能使用name来进行初始化必须使用person来
		如果结构体嵌套，外部也有name相同的，外部的name优先级比内部的优先级高
	*/
	s3 := Student2{
		Person{
			"xmaven", 19,
		},
		95.7,
	}
	fmt.Println(s3.name)
	fmt.Println(s3.print)

	fmt.Println("------------------------")
	p4 := Person{
		"booby", 18,
	}
	p4.print()

	fmt.Println("**********************************")
	p5 := Human{
		"bobby", 18, 180,
	}
	data, _ := json.Marshal(p5)
	fmt.Println(string(data))

	fmt.Println("--------------000000000000000")
	// 空结构体的使用
	// 开辟空间存储的使用比较占用空间
	var PersonSet = make(map[string]bool)
	PersonSet["booby1"] = false
	PersonSet["booby2"] = true
	PersonSet["booby3"] = false
	for key, _ := range PersonSet {
		fmt.Println(key)
	}
	fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
	p6 := Human{
		Age: 11,
	}
	data2, _ := json.Marshal(p6)
	fmt.Println(string(data2))
	fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")

	// 通过反射取值

	p7 := Human{
		"bobby", 18, 180,
	}
	t := reflect.TypeOf(p7)
	// 第一个是index 第二个是 集合里面的名字
	for _, filedName := range []string{"Name", "Age"} {
		field, ok := t.FieldByName(filedName)
		//fmt.Println(field)
		if !ok {
			panic("can not found filedName!")
		}
		fmt.Println(field.Tag.Get("bobby"))
	}
}
