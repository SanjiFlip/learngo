package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func printSlice(data []string) {
	data[0] = "java"
	for i := 0; i < 10; i++ {
		// 源数据接受不到
		data = append(data, strconv.Itoa(i))
	}
}

type slice struct {
	array unsafe.Pointer // 用来存储实际数据的数组指针，指向一块连续的内存
	len   int            // 切片中元素的数量
	cap   int            // array数组的长度
}

func main() {
	// go的切片slice 在函数参数传递的时候是值传递还是引用传递：值传递，效果呈现出了引用的效果(不完全是值传递)
	courses := []string{"go", "java", "python"}
	printSlice(courses)
	fmt.Println(courses)

	// courses2 := make([]string, 5, 10)
	// 这里必须有参数接受返回值，因为可能会触发扩容形成新的数据结构
	// courses2 = append(courses2, "abc")

	//data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//s1 := data[1:6]
	//s2 := data[2:7]
	//s2[0] = 22
	// 修改s2里面，数据s1也一样改变
	//fmt.Println(s1)
	//fmt.Println(s2)
	fmt.Println("------------------")
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := data[1:6]
	s2 := data[2:7]
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))
	// 先对s2进行放值直到扩容为止 再添加数据 成倍驼绒
	s2 = append(s2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 1, 1)
	s2[0] = 22
	// 扩容后修改s2的数据对s1 还有原数据data无改变
	// fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println("---------------------------------")
	var datanew []int
	// 512之前都是成倍扩容，后面就减缓了
	for i := 0; i < 2000; i++ {
		datanew = append(datanew, i)
		fmt.Printf("len:%d,cap:%d\r\n", len(datanew), cap(datanew))
	}
}
