package main

import "fmt"

type Person struct {
	name string
	age  int
	f    *int
}

func main() {
	// nil细节
	/*
		不同类型的数据零值不一样
		bool false
		numbers 0
		string ""
		pointer nil
		slice nil
		map nil
		channel、interface、function nil

		struct 默认值不是nil、默认值是具体字段的默认值
	*/
	p1 := Person{
		name: "xmaven",
		age:  10,
	}
	p2 := Person{
		name: "xmaven",
		age:  10,
	}
	if p1 == p2 {
		fmt.Println("yes")
	}

	// slice的默认值
	var ps []Person             // nil slice
	var ps2 = make([]Person, 0) // empty slice make 三个字段 指针数组，len ，cap
	if ps == nil {
		fmt.Println("ps is nil slice")
	}
	if ps2 == nil {
		fmt.Println("ps2 is nil slice")
	}

	// var m map[string]string //也分为nil map 和empty map
	var m map[string]string
	var m2 = make(map[string]string, 0)
	for key, value := range m {
		fmt.Println(key, value)
	}
	fmt.Println(m["name"])
	// 对未初始化的map赋值的时候会panic
	// 对初始化的就没有问题
	for key, value := range m2 {
		fmt.Println(key, value)
	}
	if m == nil {
		fmt.Println("m is nil map")
	}
	if m2 == nil {
		fmt.Println("m2 is nil map")
	}
}
