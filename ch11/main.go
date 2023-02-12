package ch11

// 单元测试 go test
// go test 命令是一个按照一定约定和组织的测试代码驱动、在包目录中，所有以_test.go后缀结尾的源码文件都会被go test执行到
// 我们写的_test.go 源码文件不用担心文件内容过多，因为go build命令不会将这些测试文件打包到最后的可执行文件中
// test文件有4类， Test开头的 功能测试 Benchmark开头的 性能测试 example 模糊测试
