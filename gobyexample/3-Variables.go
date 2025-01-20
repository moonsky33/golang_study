package main

import "fmt"

// var可以声明一个变量或多个变量
func main() {
	var a = "initial"
	// 声明字符串变量，并初始化
	fmt.Println(a)
	var b, c int = 1, 2
	// 声明多个变量，并初始化
	fmt.Println(b, c)
	var d = true
	// 声明布尔变量，初始值为false
	fmt.Println(d)
	var e int
	// 声明整数变量，初始值为0
	fmt.Println(e)
	f := "apple"
	// 短变量声明，只声明变量，不声明类型，只能在函数内部使用
	fmt.Println(f)

	excercise1()
	excercise2()
	excercise3()
}

// 小练习1：打印输出变量a的值为20，并输出为十进制、八进制、十六进制、二进制进制
func excercise1() {
	a := 20
	fmt.Printf("十进制: %d, 八进制: %o, 十六进制: %x, 二进制: %b\n", a, a, a, a)
}

// 小练习2：打印指针变量地址和指向的值,声明整型变量值为100，声明指针变量ptr指向100值，打印指针地址和指针指向的值
func excercise2() {
	var a int = 100
	var ptr *int = &a
	fmt.Println("指针地址：", ptr, "指针指向的值：", *ptr)

	value := 20
	ptr1 := &value
	fmt.Println("指针地址：", ptr1, "指针指向的值：", *ptr1)
}

// 小练习3：打印切片数量和映射变量
func excercise3() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println(slice[0:3])
	fmt.Println("slice长度：", len(slice), "slice容量:", cap(slice))
	fmt.Println("slice: %v\n", slice)

	map1 := map[string]int{"apple": 1, "banana": 2, "orange": 3}
	fmt.Println("map1的值：", map1)
	fmt.Println("map: %v\n", map1)
}
