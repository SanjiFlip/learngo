package main

import (
	"fmt"
	"strings"
)

func add(a, b int) int {
	return a + b
}

func addint32(a, b int32) int32 {
	return a + b
}

func addint64(a, b int64) int64 {
	return a + b
}

// 不兼容其他类型
func addGeneral(a, b interface{}) int {
	ai, ok := a.(int)
	if !ok {
		panic("not int type")
	}
	bi, _ := b.(int)
	return ai + bi
}

// 泛型 interface
func addGeneralAll(a, b interface{}) interface{} {
	switch a.(type) {
	case string:
		as, _ := a.(string)
		bs, _ := b.(string)
		return as + bs
	case int:
		ai, _ := a.(int)
		bi, _ := b.(int)
		return ai + bi
	case int32:
		ai, _ := a.(int32)
		bi, _ := b.(int32)
		return ai + bi
	case int64:
		ai, _ := a.(int64)
		bi, _ := b.(int64)
		return ai + bi
	case float32:
		af, _ := a.(float32)
		bf, _ := b.(float32)
		return af + bf
	case float64:
		af, _ := a.(float64)
		bf, _ := b.(float64)
		return af + bf
	default:
		panic("error type")
	}
}

func main() {
	a := 1
	b := 2
	fmt.Println(add(a, b))
	fmt.Println(addGeneral(a, b))

	fmt.Println(addGeneralAll(3, 4))
	fmt.Println(addGeneralAll(3.0, 4.0))
	re := addGeneralAll("hello", " world")
	fmt.Println(re)
	s, ok := re.(string)
	if !ok {
		panic("error type")
	}
	split := strings.Split(s, " ")
	fmt.Println(split)
}
