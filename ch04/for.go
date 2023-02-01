package main

import (
	"fmt"
	"time"
)

// go只有for循环 没有while
func main() {
	// for循环
	/*
		for init; condition: post{}
	*/
	// 相当于 while(true)
	/*for{

	}*/
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	var a int
	for {
		time.Sleep(time.Second * 2)
		fmt.Println(a)
		a++
	}

}
