package main

/*
为什么要代码规范

	1.代码规范并不是强制的，但是不同语言的一些细微的规范还是需要遵循的
	2.代码规范主要是为了方便团队内部形成一个统一的代码风格，提高代码的可读性、统一性

1.代码规范

	1.命名的规范
		包名
			1.尽量和目录保持一致
			2.尽量采取有意义的包名，简短
			3.不要和标准库名冲突
			4.包名采用全部小写
		文件名
			user_name.go 如果有多个单词可以采用蛇形命名法
		变量名
			1.蛇形：python、php
			2.驼峰：java、c、go
			userName UserName
			un string
			unad userNameAndDesc
			有一些专有名词，URLVersion
			bool类型 has、is、can、allow开头
		结构体命名
			驼峰：User
		接口命令
			接口命名基本上和结构体差不多
			接口以er结尾 / 开头用I
			type Writer interface
			type IRead interface
		常量命名
			全部大写，如果有多个单词，那么采用蛇形命名法 APP_VERSION

2.注释规范

	go提供两种注释
	1.//适合单行注释
	2.大段注释
	变量后面家注释
	包注释
	接口注释
	函数注释
	代码逻辑的注释

3.import规范

	go自带的包、第三方的包、自己内部的包
	第一组放自带的包
	第二组放第三方的包
	第三种放自己公司内部的包
*/
func main() {

}
