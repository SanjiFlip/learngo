package main

import (
	"fmt"
	// 倒入的别名
	course "learngo/ch10/user"
)

func main() {
	c := course.Course{
		Name: "go",
	}
	fmt.Println(c.Name)
	fmt.Println(course.GetCourse(c))
}
