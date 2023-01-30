package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int8 = 12
	var b uint8 = uint8(a)

	var f float32 = 3.14
	var c = int32(f)
	fmt.Println(b, c)

	var f64 = float64(a)
	fmt.Println(f64)

	// 定义别名类型
	type IT int // 类型要求严格
	var c1 IT = 13
	// var c1 IT = int(a) 错误示范
	var c2 IT = IT(a)

	fmt.Println(c1, c2)

	// 字符串转数字
	var istr = "12"
	myint, err := strconv.Atoi(istr)
	if err != nil {
		fmt.Println("convert error")
	}
	fmt.Println(myint)

	var myi = 32
	mstr := strconv.Itoa(myi)
	fmt.Println(mstr)

	// 字符串转换成 float 转换成bool
	float, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		return
	}
	fmt.Println(float)

	// 1*8的1次方+2*8的0次方
	parseInt, err := strconv.ParseInt("12", 8, 64)
	if err != nil {
		fmt.Println("ParseInt Error")
	}
	fmt.Println(parseInt)

	// parseBool 转换错误默认零值 false  1 true  0 false
	parseBool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("ParseBool Error")
	}
	fmt.Println(parseBool)

	// 基本类型转字符串
	fmt.Println("================================")
	boolStr := strconv.FormatBool(true)
	fmt.Println(boolStr)

	floatStr := strconv.FormatFloat(3.1415926, 'E', -1, 64)
	fmt.Println(floatStr)

	formatInt := strconv.FormatInt(-42, 16)
	fmt.Println(formatInt)
}
