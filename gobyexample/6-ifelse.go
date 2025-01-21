package main

import "fmt"

// go的分支结构是if-else if-else
// 并且可以进行嵌套，但是没有三元if
// 同时else等是跟在if的括号后面的
func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("8 or 7 is even")
	}
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 0 {
		fmt.Println(num, "is zero")
	} else {
		fmt.Println(num, "is positive")
	}
	exercise1()
	exercise2()
	exercise3()
}

// 小练习1：判断一个数是否是偶数
func exercise1() {
	if 6%2 == 0 {
		fmt.Println("是偶数")
	} else {
		fmt.Println("不是偶数")
	}
}

// 小练习2：找出1到100之间的完美数（完美数是指一个数等于它的因子之和，例如6=1+2+3）
func exercise2() {
	for i := 1; i <= 100; i++ {
		sum := 0
		for j := 1; j < i; j++ {
			if i%j == 0 {
				sum += j
			}
		}
		if sum == i {
			fmt.Println(i, "是完美数")
		}
	}
}

// 小练习3：根据学生分数给出对应等级评价
func exercise3() {
	fmt.Println("小练习3：根据学生分数给出对应等级评价")
	score := 85
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 && score < 90 {
		fmt.Println("B")
	} else if score >= 70 && score < 80 {
		fmt.Println("C")
	} else if score >= 60 && score < 70 {
		fmt.Println("D")
	} else {
		fmt.Println("E")

	}
}
