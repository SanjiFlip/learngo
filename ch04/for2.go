package main

import "fmt"

func main() {

	// 1-100累加
	var sum int
	for i := 1; i < 101; i++ {
		sum = sum + i
	}
	fmt.Println(sum)

	// 九九乘法口诀表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}

}
