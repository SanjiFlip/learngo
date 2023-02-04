package main

import (
	"fmt"
	"time"
)

// go中可以返回多个值
// func add(a, b int) int
func add(a int, b int) int {
	// 必须返回int值
	return a + b
}

// 函数参数传递的时候，值传递，引用传递，go语言中全部都是值传递
func add2(a, b int) (int, error) {
	return a + b, nil
}

// 省略号
func add3(items ...int) (sum int, err error) {
	for _, value := range items {
		sum = sum + value
	}
	return sum, err
}

func add4(a, b int) {
	fmt.Printf("sum is %d\r\n", a+b)
}

func runForever() {
	for {
		time.Sleep(time.Second)
		fmt.Println("doing...")
	}
}

func cal(op string, items ...int) func() {
	switch op {
	case "+":
		return func() {
			fmt.Println("这是加法")
		}
	case "-":
		return func() {
			fmt.Println("这是减法")
		}
	default:
		return func() {
			fmt.Println("既不是加法又不是减法")
		}
	}
}

func cal2(myfunc func(items ...int) int) int {
	return myfunc()
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}

var local int

func autoIncrement() int {
	local += 1
	return local
}

func autoIncrement2() func() int {
	local := 0 // 一个函数中访问另一个函数的局部变量 不行的
	return func() int {
		local += 1
		return local
	}
}

func main() {
	// go函数支持普通函数、匿名函数、闭包
	/*
		go种函数是"一等公民"
			1.函数本身可以当作变量
			2.匿名函数 闭包
		 	3.函数可以满足接口
	*/
	fmt.Println(add(1, 2))
	sum, _ := add2(1, 2)
	fmt.Println(sum)
	sum2, _ := add3(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(sum2)

	// 第一种函数当作变量
	funcVar := add
	fmt.Println(funcVar(1, 2))
	cal("+")()
	sum3 := cal2(func(items ...int) int {
		sum := 0
		for _, value := range items {
			sum += value
		}
		return sum
	})
	fmt.Println(sum3)
	fmt.Println("-------------------")
	callback(1, add4)

	// 匿名函数
	callback(1, func(a, b int) {
		fmt.Printf("total is %d\r\n", a+b)
	})

	locaFunc := func(a, b int) {
		fmt.Printf("total is %d\r\n", a+b)
	}
	callback(2, locaFunc)

	fmt.Println("=================")
	// 闭包
	// 有个需求，我希望有个函数每次调用返回的结果值都是增加一次之后的值

	for i := 0; i < 5; i++ {
		fmt.Println(autoIncrement())
	}
	fmt.Println("*****************")
	nextFunc := autoIncrement2()
	for i := 0; i < 5; i++ {
		fmt.Println(nextFunc())
	}
	// 重新定义会归零
	nextFunc2 := autoIncrement2()
	for i := 0; i < 3; i++ {
		fmt.Println(nextFunc2())
	}
}
