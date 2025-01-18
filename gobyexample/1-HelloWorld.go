package main

import "fmt"

// 这是go的入口包声明，所有go程序都需主包声明为main，也是告诉编译器是该程序的起点

// // 引入fmt包，该包也就是提供格式化输入输出函数
func main() {
	fmt.Println("Hello World,This my first go study-test,today is 2025-1-18! go!")
	table()
}

// func main()是程序的入口函数，main函数是程序的入口函数，程序从main函数开始执行，main函数执行完毕程序结束
// 基本上god代码都得被包裹在里面,println和python中的print都是输入输出打印

// 你也可以如下,声明一个变量为什么什么，来打印该变量。
// ：=是短变量并赋值的操作符，会自动推断并进行复制
// func main() {
// 	message := "Hello World,This my first go study-test,today is 2025-1-18! go!"
// 	fmt.Printf(message)
// }

// 现在来一个练习题，一个打印表格

func table() {
	fmt.Printf("|%s\t|%s\t|%s\t|\n", "姓名", "年龄", "成绩")
	fmt.Printf("|%s\t|%d\t|%d\t|\n", "张三", 18, 100)
	fmt.Printf("|%s\t|%d\t|%d\t|\n", "李四", 19, 99)
}

// 这里就发现了一些制表符相关符号也是可以用的，并且你要调用一个新的函数调用，就得在main函数中添加该函数也就是table()，就是所调用的
