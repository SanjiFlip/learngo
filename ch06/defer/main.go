package main

import "fmt"

func main() {
	// 连接数据库、打开文件、开始锁、无论如何 最后都要记得去关闭数据库、关闭文件、释放锁

	// 相当于Java的finally defer后面的代码会放在函数return之前执行 可以有多个defer
	// 多个defer相当于栈的概念，后进先出
	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("main")
	return
	/*

	 */

}
