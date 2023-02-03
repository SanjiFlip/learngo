package main

import (
	"container/list"
	"fmt"
)

func main() {
	// list 相当于链表
	// 空间不连续，空间浪费，性能差异大 插入元素简单 频繁插入简单 查询麻烦，需要遍历
	// 工作种使用最多的就是slice、map

	var mylist list.List

	// 放在尾部
	mylist.PushBack("go")
	mylist.PushBack("grpc")
	mylist.PushBack("mysql")
	fmt.Println(mylist)
	// 遍历打印值，正序
	for i := mylist.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	fmt.Println("----------")
	// 反向遍历，逆序
	for i := mylist.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}

	fmt.Println("---------------")
	mylist2 := list.New()
	// 尾部存放数据
	mylist2.PushBack("go")
	mylist2.PushBack("grpc")
	mylist2.PushBack("mysql")
	//头部存放数据
	mylist2.PushFront("gin")

	i := mylist2.Front()
	for ; i != nil; i = i.Next() {
		if i.Value.(string) == "grpc" {
			break
		}
	}
	// 必须找到插入的element
	mylist2.InsertBefore("java", i)
	// 删除元素也是需要先找到相关的element
	mylist2.Remove(i)
	for i := mylist2.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	/*
		集合类型4种
		1. 数组- 不同长度的数组类型不一样
		2. 切片- 动态数组，用起来方便，性能也很高，我们要尽量使用这种
		3. map 常用
		4. list:用的少
	*/
}
