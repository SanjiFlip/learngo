package main

import "fmt"

func main() {

	// go语言提供了那些几盒类型的数据结构，数组、切片(slice)、mymap、list
	// 数组 定义 var name [count]int
	var courses1 [3]string // courses类型 ，数组 只有三个元素的数组
	var courses2 [4]string
	// []string 和 [3]string 是两种不同的类型 前面一种是切片 后面一种是数组
	// course1 = courses2 // 错误的
	courses1[0] = "go"
	courses1[1] = "grpc"
	courses1[2] = "gin"

	fmt.Printf("%T %T\r\n", courses1, courses2)
	fmt.Println(courses1)

	for _, value := range courses1 {
		fmt.Println(value)
	}

}
