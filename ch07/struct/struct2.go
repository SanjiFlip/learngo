package main

import "fmt"

type Person2 struct {
	name    string
	age     int
	address string
	height  float32
}

func main() {
	var p Person2
	p.age = 20

	fmt.Println(p.name)

	// 匿名结构体，匿名函数
	address := struct {
		province string
		city     string
		address  string
	}{
		"北京市",
		"通州区",
		"xxx",
	}
	fmt.Println(address.city)
}
