package main

import "fmt"

// 全局变量和局部变量
// 第一种定义方式
//var name = "xmaven"
//var age = 18
//var ok bool

// 第二种定义方式
var (
	name = "xmaven"
	age  = 18
	ok   bool
)

func main() {
	//go是静态语言，静态语言和和动态语言相比变量差异很大
	//1. 变量必须先定义后使用 2.变量必须有类型 3. 类型定下无法改变
	// 定义变量的方式
	//var name int
	//name = 1

	fmt.Println(name)

	//var age = 1
	age := 1

	var age2 int

	// go语言种变量定义了不使用是不行的
	fmt.Println(age)

	fmt.Println(age2)

	// 2.多变量定义
	var user1, user2, user3 = "xmaven1", 1, "xmaven3"
	fmt.Println(user1, user2, user3)

	/*
		注意:
			变量必须先定义才能使用
			go语言是静态语言，要求变量的类型和赋值类型一致
			变量名不能冲突
			简洁变量不能用于全局 例如age2 := 11
			变量是有零值的,也就是默认值
			定义了变量一定要使用，否者会报错
	*/
}
