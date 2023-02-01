package main

import "fmt"

func main() {
	//for 循环还有一种用法 for range 主要对字符串、数组、切片、map、channel
	/*
		for key,value := range {
		}
	*/
	name := "xmaven go学习之路"
	print(len(name))
	for index, value := range name {
		//fmt.Println(index, value)
		fmt.Printf("%d,%c\r\n", index, value)
	}
	// _匿名变量 不能去掉_
	for _, value := range name {
		fmt.Printf("%c\r\n", value)
	}
	/*
		字符串 字符串的索引(key) 字符串对应的索引的字符值的拷贝(value)       如果不写key，返回的就是索引
		数组	    数组的索引	    索引对应的值的拷贝	                    如果不写 key，那么返回的是索引
		切片	    切片的索引	    索引对应的值的拷贝	                    如果不写 key，那么返回的是索引
		map	    map 的 key	    value 返回的是 key 对应的值的拷贝	    如果不写 key，那么返回的是 map 的值
		channel		            value 返回的是 channel 接受的数据
	*/
	fmt.Println("-----------------------")

	for index := range name {
		// 如果name中有中文打印出来会乱码
		fmt.Printf("%d %c\r\n", index, name[index])
	}
	fmt.Println("-----------------------")
	// 如果非要打印中文可以这样操作
	nameRune := []rune(name)
	for index := range nameRune {
		fmt.Printf("%d,%c\r\n", index, nameRune[index])
	}

	fmt.Println("1111111111111111111111111")

	for i := 0; i < len(nameRune); i++ {
		fmt.Printf("%d,%c\r\n", i, nameRune[i])
	}
}
