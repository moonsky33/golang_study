package main

import (
	"fmt"
	"time"
)

// for是go唯一的循环结构，后面跟着的就是循环条件，后面没跟任何就是无限循环，可以用break跳出
func main() {
	// 初始化变量i为1，用于控制循环的次数
	i := 1

	// 循环打印变量i的值，直到i大于3为止
	for i <= 3 {
		// 打印当前i的值
		fmt.Println(i)
		// 增加i的值，以避免无限循环
		i = i + 1
	}
	// j初始变量为0，只要小于3就循环，每次j+1，直到j大于3为止
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	// 循环打印range的值，range是go的关键字，用于遍历数组、切片、字符串、map、通道等数据结构
	for i := range 3 {
		fmt.Println("range", i)
	}
	// 无限循环，可以用break跳出
	for {
		fmt.Println("loop")
		break
	}
	// 循环打印n的值，只要n是偶数就跳过，直到n大于5为止
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	iterativechannel()
	// goto_test()
	// mapSlice([]int{1, 2, 3, 4, 5}, func(n int) int { return n * 2 })
	// map_test()
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
}

// for语法
// for initialization;condition;post{循环语句}
// initialization:初始化语句，只执行一次，可以省略，如果省略，则必须在循环体内定义变量。
// condition:条件表达式，循环条件，如果为false，则循环结束。
// post:循环体执行完之后执行的语句，可以省略。

// 拓展1：迭代通道

// 迭代通道并不是特定数据类型，而是指使用for-->range循环遍历通道。
// 能够持续从通道接收数据，直到通道关闭，并且无需每次接收手动检查通道是否关闭
// 确保了迭代通道保证发送的数据都被接收直到退出循环。保证数据完整性
// 并且提高代码可读性，更简洁明了
func iterativechannel() {
	// 创建一个通道ch，并向其中发送三个值
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()
	for value := range ch {
		fmt.Println(value)
	}
	time.Sleep(time.Second)
}

// 拓展二：标签和goto

// 请注意：过度使用goto会让代码难以理解和维护

// func goto_test() {
// 	// 当使用goto时，应确保goto跳转的位置是唯一的，并且不会导致循环嵌套。
// 	// 当i等于1且j等于1时，跳转到outerloop标签处，跳过循环体。
// outerloop:
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 3; j++ {
// 			if i == 1 && j == 1 {
// 				goto outerloop
// 			}
// 			fmt.Printf("i: %d, j: %d\n", i, j)
// 		}
// 	}
// }

// 拓展三：for循环的函数式编程风格
// 可以使用for循环来完成函数式编程擦欧总，例如map（映射）、filter（过滤）、reduce（归约）

// 使用mapSlice函数接收一个切片和一个函数mapper，使用for循环将mapper函数应用到切片的每个元素上，生成一个新的切片
// func mapSlice(nums []int, mapper func(int) int) []int {
// 	result := make([]int, len(nums))
// 	for i, v := range nums {
// 		result[i] = mapper(v)
// 	}
// 	return result
// }

// func map_test() {
// 	nums := []int{1, 2, 3, 4, 5}
// 	mapped := mapSlice(nums, func(n int) int {
// 		return n * 2
// 	})
// 	fmt.Println(mapped)
// }

// 使用filterSlice函数接收一个切片和一个函数predicate，使用for循环将predicate函数应用到切片的每个元素上，生成一个新的切片，只包含predicate函数返回true的元素

// func filterSlice(nums []int, predicate func(int) bool) []int {
// 	result := []int{}
// 	for _, v := range nums {
// 		if predicate(v) {
// 			result = append(result, v)
// 		}
// 	}
// 	return result
// }

// func filter_test() {
// 	nums := []int{1, 2, 3, 4, 5}
// 	filtered := filterSlice(nums, func(n int) bool {
// 		return n%2 == 0
// 	})
// 	fmt.Println(filtered)
// }

// 使用reduceSlice函数接收一个切片和一个函数reducer，使用for循环将reducer函数应用到切片的每个元素上，生成一个新的切片，只包含reducer函数返回true的元素
// func reduceSlice(nums []int, reducer func(int, int) int) int {
// 	result := 0
// 	for _, v := range nums {
// 		result = reducer(result, v)
// 	}
// 	return result
// }

// func reduce_test() {
// 	nums := []int{1, 2, 3, 4, 5}
// 	reduced := reduceSlice(nums, func(acc, n int) int {
// 		return acc + n
// 	})
// 	fmt.Println(reduced)
// }

// 拓展四：for循环与并发\
// 可以在for循环中使用并发机制，例如使用goroutine和channel，以实现并发控制。

// func concurrent_test() {
// 	var wg sync.WaitGroup
// 	nums := []int{1, 2, 3, 4, 5}
// 	for _, v := range nums {
// 		wg.Add(1)
// 		go func(n int) {
// 			defer wg.Done()
// 			fmt.Println(n)
// 		}(v)
// 	}
// 	wg.Wait()
// }

// 小练习1：使用for循环打印从1到10的数字
func exercise1() {
	fmt.Println("小练习1：打印1到10数字")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// 小练习2：计算1-100整数和，并打印
func exercise2() {
	fmt.Println("小练习2：计算1-100整数和")
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1-100整数和为：", sum)
}

// 小练习3：打印九九乘法表
func exercise3() {
	fmt.Println("小练习3：打印九九乘法表")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Printf("%d x %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
}

// 小练习4：斐波那契数列，前20项
// f(o)=0,f(1)=1,f(n)=f(n-1)+f(n-2)
func exercise4() {
	a, b := 0, 1
	for i := 0; i < 20; i++ {
		fmt.Println("第", i, "项", a)
		a, b = b, a+b
	}
}

// 小练习5：计算阶乘
// n!=n*(n-1)*(n-2)*...*1

func exercise5() {
	n := 5
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	fmt.Println("5的阶乘：", factorial)
}
