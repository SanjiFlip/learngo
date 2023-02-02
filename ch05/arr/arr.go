package main

import "fmt"

func main() {

	//数组的初始化1
	//var courses1= [3]string{"go", "grpc", "gin"}
	//courses1 := [3]string{"go", "grpc", "gin"}
	var courses1 [3]string = [3]string{"go", "grpc", "gin"}

	for _, value := range courses1 {
		fmt.Println(value)
	}

	// 初始化2
	courses2 := [3]string{2: "gin"}
	for _, value := range courses2 {
		fmt.Println(value)
	}

	// 初始化3 省略号就是变长
	courses3 := [...]string{"go", "grpc"}
	for _, value := range courses3 {
		fmt.Println(value)
	}

	for i := 0; i < len(courses3); i++ {
		fmt.Println(courses3[i])
	}
	fmt.Println("-----------------------")
	courses4 := [...]string{"go", "gin"}
	if courses3 == courses4 { // 比较里面每个index下的value
		fmt.Println("equal")
	}
	// 多维数组
	var courseInfo [3][4]string
	courseInfo[0] = [4]string{"java", "1h", "xmaven", "qq"}
	courseInfo[1][0] = "go"
	courseInfo[1][1] = "2h"
	courseInfo[1][2] = "tom"
	courseInfo[1][3] = "web"
	courseInfo[2] = [4]string{"python", "1.5h", "jacklove", "douyin"}
	fmt.Println(len(courseInfo))
	for i := 0; i < len(courseInfo); i++ {
		for j := 0; j < len(courseInfo[i]); j++ {
			fmt.Printf("%s\t", courseInfo[i][j])
		}
		fmt.Println()
	}
	for _, row := range courseInfo {
		fmt.Println(row)
	}
	for _, rowArr := range courseInfo {
		for _, column := range rowArr {
			fmt.Printf("%s\t", column)
		}
		fmt.Println()
	}
}
