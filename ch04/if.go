package main

import "fmt"

/*
if 布尔表达式{
	逻辑
}
*/

func main() {
	// if条件判断
	age := 22
	country := "美国"
	// 可以使用括号进行分组
	if age < 18 {
		if country == "中国" {
			fmt.Println("未成年")
		}
	} else if age == 18 {
		fmt.Println("刚刚成年")
	} else {
		fmt.Println("成年")
	}

	if age < 18 {
		fmt.Println("未成年")
	}
	if age == 18 {
		fmt.Println("刚好成年")
	}
	if age > 18 {
		fmt.Println("未成年")
	}
}
