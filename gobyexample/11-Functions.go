package main

import "fmt"

// go的函数和python什么的都差不多的意思

// 定义plus函数为a+b。获取整数型a和b参数
func plus(a int, b int) int {
	return a + b
}

//同理
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	// 调用plus函数，传入参数1和2，赋值给res
	fmt.Println("1+2=", res)
	res = plusPlus(1, 2, 3)
	// 同上
	fmt.Println("1+2+3=", res)
}
