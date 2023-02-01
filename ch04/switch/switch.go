package main

import "fmt"

func main() {
	/*
		switch var1 {
		    case val2:
		        ...
		    case val3:
		        ...
		    default:
		        ...
	*/

	// 中文的星期几 输出对应的英文
	day := "星期三"
	switch day {
	case "星期一":
		fmt.Println("Monday")
	case "星期二":
		fmt.Println("Tuesday")
	case "星期三":
		fmt.Println("Wednesday")
	case "星期四":
		fmt.Println("Thursday")
	case "星期五":
		fmt.Println("Friday")
	case "星期六":
		fmt.Println("Saturday")
	case "星期天":
		fmt.Println("Sunday")
	default:
		fmt.Println("Unknow")
	}
	score := 60
	switch {
	case score < 60:
		fmt.Println("D")
	case score >= 60 && score < 70:
		fmt.Println("C")
	case score >= 70 && score < 80:
		fmt.Println("B")
	case score >= 80 && score < 90:
		fmt.Println("A")
	case score >= 90 && score <= 100:
		fmt.Println("A+")
	default:
		fmt.Println("SSS+")
	}
	score2 := 70
	switch score2 {
	case 60, 70, 80:
		fmt.Println("Wow")
	default:
		fmt.Println("default")
	}
}
