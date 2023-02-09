package main

import "fmt"

// 接口的定义
type Duck interface {
	// 方法的声明
	Gaga()
	Walk()
	Swimming()
}

type pskDuck struct {
	legs int
}

func (pd *pskDuck) Gaga() {
	fmt.Println("嘎嘎")
}

func (pd *pskDuck) Walk() {
	fmt.Println("this is pskDuck walking")
}

func (pd *pskDuck) Swimming() {
	fmt.Println("this is pskDuck swimming")
}

func main() {
	// go语言的接口，鸭子类型，php、python
	// go语言中处处都是interface，到处都是鸭子类型 duck typing
	/*
		当看到一只鸟走起来像鸭子、有用起来像鸭子、叫起来也像鸭子，那么这只鸟就是鸭子
		动词，方法，具备某些方法
	*/
	var d Duck = &pskDuck{}
	d.Walk()
}
