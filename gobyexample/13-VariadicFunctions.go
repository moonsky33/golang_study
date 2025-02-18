package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	// 计算sum的参数个数总和
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	// 调用sum函数，传入两个函数
	sum(1, 2, 3)
	num := []int{1, 2, 3, 4}
	// 将nums切片里面所有传入sum函数
	sum(num...)
}
