package main

import "fmt"

func main() {
	// 运算符 + - * / % ++ --
	var a, b = 1, 2
	var astr, bstr = "hello", "xmaven"
	fmt.Println(a + b)
	fmt.Println(astr + bstr)
	// % 取余
	fmt.Println(3 % 2)
	//a++
	//a += 1
	//a = a + 1

	a--
	a -= 1
	a = a - 1
	fmt.Println(a)

	// 逻辑运算符
	var abool = true
	var bbool = false
	if abool && bbool {
		fmt.Println("全真为真")
	}
	if abool || bbool {
		fmt.Println("有真为真")
	}
	if !abool {
		fmt.Println("取反")
	}

	// 位运算符,性能要求高的时候会考虑与运算
	var A = 60
	var B = 13
	fmt.Println(A & B)
	fmt.Println(A | B)
	fmt.Println(A ^ B)

}
