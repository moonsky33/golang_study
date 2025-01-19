package main

import (
	"fmt"
)

func main() {
	fmt.Println("go" + "lang")
	// 字符串拼接
	fmt.Println("1+1=", 1+1)
	// 整数相加
	fmt.Println("7.0/3.0=", 7.0/3.0)
	// 浮点数相除
	fmt.Println(true && false)
	// 逻辑与，两边都是true才是true否则false
	fmt.Println(true || false)
	// 逻辑或，两边只要有一个true就是true,否则false
	fmt.Println(!true)
	// 逻辑非，取反，即true变false，false变true
	exercise1()
	// 调用练习1函数
	exercise2()
	// 调用练习2
	exercise3()
	// 计算机原周长和面积
}

// 小练习1：将两个字符串拼接起来，并打印出来。
func exercise1() {
	str1 := "hello"
	str2 := "world"
	fmt.Println("拼接字符串是：", str1+str2)
}

// 小练习2：多个不同类型数据组合打印一个字符串
func exercise2() {
	name := "小明"
	age := 18
	height := 1.75
	fmt.Printf("我的名字是%s,我今年%d，身高%.2f米\n", name, age, height)
	// %s表示字符串占位符，%d表示整数占位符，%.2f表示保留两位小数的浮点数占位符，后面跟着会进行顺序替换
}

// 小练习3：计算圆面积和周长，半径为2.0
func exercise3() {
	r := 2.0
	pai := 3.14
	c := 2 * pai * r
	// 2 * pai * r 表示周长公式
	s := pai * r * r
	fmt.Printf("圆的周长是%.2f,面积是%.2f\n", c, s)
}
