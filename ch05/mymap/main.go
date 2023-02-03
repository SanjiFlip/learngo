package main

import "fmt"

func main() {
	// map 是一个key(索引)和value(值)的无序集合，主要查询方便时间复杂度O(1)
	var courseMap = map[string]string{
		"go":   "go工程师",
		"grpc": "grpc入门",
		"gin":  "gin深入理解",
	}
	// 取值
	fmt.Println(courseMap["grpc"])

	// 放值
	courseMap["mysql"] = "mysql原理"
	fmt.Println(courseMap["mysql"])

	/*
		// 错误示范 不能给nil(空) map放值
		var courseMap2 map[string]string // nil， 想要设置值，必须要初始化
		courseMap2["mysql"] = "mysql原理"
		fmt.Println(courseMap2["mysql"])
	*/
	var courseMap2 = map[string]string{} //想要设置值，必须要初始化 ,可以什么都不放
	courseMap2["mysql"] = "mysql原理2"
	fmt.Println(courseMap2["mysql"])

	// make 是内置函数，主要用于初始化slice ，map ，channel
	var courseMap3 = make(map[string]string, 3)
	courseMap3["mysql"] = "mysql原理3"
	fmt.Println(courseMap3["mysql"])

	/*
		map 必须初始化才能使用
		1.var courseMap2 = map[string]string{}
		2.var courseMap3 = make(map[string]string, 3)
		但是slice可以不初始化
		map 赋值
		1.初始化时赋值
		2.初始化完成后添加值
	*/
	//var m []string
	//if m == nil {
	//	fmt.Println("yes")
	//}
	//m = append(m, "a")
	fmt.Println("--------------------")
	// 遍历
	// map 的key是数组 基础类型 不能是切片类型
	for key, value := range courseMap {
		fmt.Println(key, value)
	}
	for key := range courseMap {
		fmt.Println(key, courseMap[key])
	}
	// map 是无序的，而且不保证每次打印都是一致的顺序
	courseMap["java"] = ""
	d, ok := courseMap["java"]
	// ok true 找的到，false找不到
	if !ok {
		fmt.Println("not in")
	} else {
		fmt.Println("find ", d)
	}

	if d, ok := courseMap["es"]; !ok {
		fmt.Println("not in")
	} else {
		fmt.Println("find ", d)
	}

	fmt.Println("------------------")
	//删除元素
	delete(courseMap, "grpc")
	// 删除一个不存在的元素也不会报错
	delete(courseMap, "ggggg")

	//很重要的提示，map不是线程安全的
	fmt.Println(courseMap)

}
