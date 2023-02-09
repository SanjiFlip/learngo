package main

import "fmt"

type MyWriter interface {
	Write(string) error
}

type MyCloser interface {
	Close() error
}

type writerCloser struct {
	// 可以显式也可以隐式 mw MyWriter
	MyWriter // interface也是一个类型 想放入一个写入文件的实现，放入一个写数据库的实现
}

//func (wc *writerCloser) Write(string) error {
//	fmt.Println("write string")
//	return nil
//}
//func (wc *writerCloser) Close() error {
//	fmt.Println("close")
//	return nil
//}

type fileWriter struct {
	filePath string
}

type databaseWriter struct {
	host string
	port string
	db   string
}

func (wc *databaseWriter) Write(string) error {
	fmt.Println("write sth to database")
	return nil
}
func (wc *fileWriter) Write(string) error {
	fmt.Println("write string to file")
	return nil
}
func (wc *fileWriter) Close() error {
	fmt.Println("close")
	return nil
}

func main() {
	//var mw MyWriter = &writerCloser{}
	//var mc MyCloser = &writerCloser{}

	var mw MyWriter = &writerCloser{
		&fileWriter{},
	}
	mw.Write("xmaven")

	var mw2 MyWriter = &writerCloser{
		&databaseWriter{},
	}
	mw2.Write("xmaven")
}
