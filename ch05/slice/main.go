package main

import "fmt"

func main() {
	// go折中，slice 切片 - 数组
	var courses []string
	fmt.Printf("%T\r\n", courses)

	// 这个方法很特别 ,第一次用很头疼，原理
	courses = append(courses, "go")
	courses = append(courses, "grpc")
	courses = append(courses, "gin")
	fmt.Println(courses[1])

	// 切片的初始化 3种方式 1：从数组直接创建 2：使用stirng{} 3：make  常用的也就后两种
	allCourses := [5]string{"go", "grpc", "gin", "mysql", "es"}
	courseSlice := allCourses[0:len(allCourses)] // 左闭右开 [) python的语法
	fmt.Println(courseSlice)

	courseSlice2 := []string{"go", "grpc", "gin", "mysql", "es"}
	courseSlice3 := courseSlice2[0:]
	fmt.Println(courseSlice3)

	fmt.Println("==============================")

	// make 函数
	allCourses2 := make([]string, 3)
	allCourses2[0] = "c"
	allCourses2[1] = "c"
	allCourses2[2] = "c"
	//allCourses2[3] = "c" 超过长度只能使用append ，没有默认给初始存储空间不能存储元素
	allCourses2 = append(allCourses2, "d")
	fmt.Println(allCourses2)

	// 下面操作只能使用append 不能直接赋值
	//var allCourses3 []string
	//allCourses3[0] = "c"

	fmt.Println("==============================")
	// 访问切片的元素,访问单个，访问多个
	courseSlice5 := []string{"go", "grpc", "gin", "mysql", "es"}
	fmt.Println(courseSlice5[2])
	// [start:end] start从零开始，end结束 同样是左开右闭原则
	// 1。如果只有个start开始 没有end ，就是从start到后面所有的
	// 2。如果只用end 没有start 就是从第一个元素到end截止不含end 同样是左开右闭原则
	// 3。如果没有start 没有end 就是打印全部元素
	// 4。如果有start 有end ，就是冲start开始到end  同样是左开右闭原则
	fmt.Println(courseSlice5[1:4])
	fmt.Println(courseSlice5[1:])
	fmt.Println(courseSlice5[:3])
	fmt.Println(courseSlice5[:])

	fmt.Println("==============================")

	courseSlice6 := []string{"go", "grpc"}
	courseSlice6 = append(courseSlice6, "gin")
	fmt.Println(courseSlice6)

	fmt.Println("==============================")
	// 合并两个slice
	courseSlice7 := []string{"mysql", "python", "java"}
	courseSlice8 := []string{"go", "grpc"}
	// 第一种合并方法
	//for _, value := range courseSlice8 {
	//	courseSlice7 = append(courseSlice7, value)
	//}
	// 第二种合并方法 后面... 可以把切片里面的元素打散 也可以放这种切片 courseSlice8[1:]...
	courseSlice7 = append(courseSlice7, courseSlice8...)

	fmt.Println(courseSlice7)

	fmt.Println("==============================")

	// 如何删除元素 比较麻烦 比如删除 mysql
	courseSlice9 := []string{"go", "grpc", "mysql", "es", "gin"}
	myslice := append(courseSlice9[:2], courseSlice9[3:]...)
	fmt.Println(myslice)

	courseSlice9 = courseSlice9[:3]
	fmt.Println(courseSlice9)

	fmt.Println("==============================")
	courseSlice10 := []string{"go", "grpc", "mysql", "es", "gin"}
	//复制slice
	courseSliceCopy := courseSlice10
	courseSliceCopy2 := courseSlice10[:]
	fmt.Println(courseSliceCopy, courseSliceCopy2)
	// 拷贝slice

	// var courseSliceCopyNew []string  //错误定义 这种空间不回拷贝，错误用法
	var courseSliceCopyNew = make([]string, len(courseSlice10))
	copy(courseSliceCopyNew, courseSlice10)
	fmt.Println(courseSliceCopyNew)
}
