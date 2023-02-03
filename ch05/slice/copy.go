package main

import "fmt"

func main() {
	courseSlice := []string{"go", "grpc", "mysql", "es", "gin"}
	//复制slice
	// courseSliceCopy1 := courseSlice
	// 原数据改变会影响拷贝的数据
	courseSliceCopy2 := courseSlice[:]
	fmt.Println(courseSliceCopy2)
	// 拷贝slice
	// var courseSliceCopyNew []string  //错误定义 这种空间不会拷贝，错误用法
	// 下面这种拷贝方式开辟了新的空间，原数据改变，拷贝后的数据不会改变
	var courseSliceCopy3 = make([]string, len(courseSlice))
	copy(courseSliceCopy3, courseSlice)
	fmt.Println(courseSliceCopy3)

	fmt.Println("--------------------")
	courseSlice[0] = "java"
	fmt.Println(courseSliceCopy2)
	fmt.Println(courseSliceCopy3)
}
