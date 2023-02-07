package main

import "fmt"

// 定义结构体
// 存放多个person的信息到一个集合中去
type Person struct {
	name    string
	age     int
	address string
	height  float32
}

func main() {

	// 如何初始化结构体
	p1 := Person{"xmaven", 18, "武汉", 1.99}
	// p2 中可以指定赋值
	p2 := Person{
		name:    "xmaven2",
		age:     11,
		address: "北京",
		height:  1.78,
	}
	var persons []Person
	persons = append(persons, p1)
	persons = append(persons, Person{
		name: "去你妈的",
	})
	persons2 := []Person{
		{
			name: "bbbb",
		},
		{
			name:    "xmaven2",
			age:     11,
			address: "北京",
			height:  1.78,
		},
		{
			age: 19,
		},
	}
	fmt.Println(p1, p2)
	fmt.Println(persons)
	fmt.Println(persons2)
}
