package main

import "fmt"

func deferReturn() (ret int) {
	defer func() {
		ret++
	}()
	return 10
}

// error , panic , recover
// go语言错误处理的理念，一个函数kennel出错，try catch 去包住这个函数
// 开发函数的人需要有一个返回值告诉调用者是否成功，go设计者要求我们必须处理err，代码中会出现大量的if err!=nil
// go设计者认为必须要处理这个error，防御编程

func main() {
	// defer 有能力修改返回值的
	ret := deferReturn()
	fmt.Printf("ret = %d\r\n", ret)
}
