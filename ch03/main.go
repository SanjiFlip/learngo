package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 长度计算
	// 如果你想知道一个字符串(中文)长度 ,如果你只有英文，你可以直接用len
	ename := "xmaven"
	fmt.Println(len(ename))
	name := "xmave笔记之路"
	bytes := []rune(name)
	fmt.Println(len(bytes))

	//转义符
	// 或者使用 `Go"学习之路"`
	courseName := "Go\"学习之路\""
	fmt.Println(courseName)
	courseName2 := "\tGo学习之路"
	fmt.Println(courseName2)
	courseName3 := "Hello\r\nGo学习之路"
	fmt.Println(courseName3)
	fmt.Print("hello\r\n")
	fmt.Print("world")

	//格式化输出
	fmt.Println("============================")
	username := "xmaven"
	age := 18
	address := "武汉"
	mobile := "1008611"
	//极难维护
	fmt.Println("用户名: "+username, ","+"年龄: "+strconv.Itoa(age), ","+"地址:"+address, ","+"电话号码: "+mobile)
	//这个非常常用 但是性能没有上面的好
	fmt.Printf("用户名:%s, 年龄:%d, 地址:%s, 电话号码:%s\r\n", username, age, address, mobile)
	userMsg := fmt.Sprintf("用户名:%s, 年龄:%d, 地址:%s, 电话号码:%s\r\n", username, age, address, mobile)
	fmt.Println(userMsg)

	//通过string的builder进行字符串拼接，高性能
	var builder strings.Builder
	builder.WriteString("用户名:")
	builder.WriteString(username)
	builder.WriteString(",年龄:")
	builder.WriteString(strconv.Itoa(age))
	builder.WriteString(",地址:")
	builder.WriteString(address)
	builder.WriteString(",电话号码:")
	builder.WriteString(mobile)

	fmt.Println(builder.String())
}
