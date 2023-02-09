package main

import "fmt"

func mprint(datas ...interface{}) {
	for _, value := range datas {
		fmt.Println(value)
	}
}

func mprint2(data interface{}) {
	fmt.Println(data)
}

type myinfo struct {
}

func (mi *myinfo) Error() string {
	return "我不是error"
}

func main() {
	var data = []interface{}{
		"booby", 18, 180,
	}
	// ...参数打散，任意参数都可以
	mprint(data...)
	// 具体参数传入变长参数函数，会报错
	var data2 = []string{
		"booby", "booby2", "booby3",
	}
	// 直接放数据是不行的 需要迂回放置数据
	// mprint(data2...)
	mprint2(data2)
	var datai = []interface{}{}
	for _, value := range data2 {
		datai = append(datai, value)
	}
	mprint(datai)

	var err error
	err = &myinfo{}
	fmt.Println(err)
}
