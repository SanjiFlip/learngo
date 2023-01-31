package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串比较
	a := "hello"
	b := "bello"
	fmt.Println(a != b)

	// 字符串大小的比较
	fmt.Println(a > b)

	// 是否包含
	name := "imooc体系课go开发工程师go"
	fmt.Println(strings.Contains(name, "go"))

	// 字符的出现次数
	fmt.Println(strings.Count(name, "0"))

	//分割字符串 数组
	fmt.Println(strings.Split(name, "-"))

	//字符串是否包含前缀，是否包含后缀
	fmt.Println(strings.HasPrefix(name, "imooc"))

	fmt.Println(strings.HasSuffix(name, "工程师"))

	// 查询字串出现的位置
	fmt.Println(strings.Index(name, "go"))
	fmt.Println(strings.IndexRune(name, []rune(name)[8]))

	// 子串替换 最后一个参数替换的次数，若为复数则全部替换
	fmt.Println(strings.Replace(name, "go", "java", -1))

	//大小写转换
	fmt.Println(strings.ToLower("GO"))
	fmt.Println(strings.ToUpper("java"))

	// 去掉特殊字符 只能去掉左右两边的字符 中间的无法去除 cutset 可以写字符串只要左右两边在cutset里面的全部去除
	fmt.Println(strings.Trim(" Hello go", " "))

}
