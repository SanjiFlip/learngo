package main

import (
	"fmt"
	_ "learngo/ch09/user" // 自动调用init方法
	. "learngo/ch10/user" // 这种用法尽量少用 一般都是使用定义别名
)

func main() {
	c := Course{Name: "go"}
	fmt.Println(c.Name)
	fmt.Println(GetCourse(c))
}
