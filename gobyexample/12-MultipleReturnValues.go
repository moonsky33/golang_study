package main

import "fmt"

// vals函数里接入两个值，默认分别返回3，7
func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	// 赋值，下面两打印
	fmt.Println(a)
	fmt.Println(b)
	_, c := vals()
	// 第一个赋值空白标识符忽略不需要的返回值，第二个赋值c，下面打印c
	fmt.Println(c)
}
