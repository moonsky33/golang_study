package main

import "fmt"

// go语言中的数组具有固定的长度，其长度在声明时确定且不能更改
// 可以通过索引访问，索引从0开始
// ...让编译器推断长度和使用索引初始化元素

// 注：如果访问数组元素时，索引超出数组范围，会导致运行报错（越界）
func main() {

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
	exercise1()
	exercise2()
	exercise3()
}

// 小练习1：计算数组的平均值
func exercise1() {
	array := [5]int{1, 2, 3, 4, 5}
	sum := 0
	for _, v := range array {
		sum += v
	}
	fmt.Println("avg:", sum/len(array))
}

// 小练习2：找出数组出现次数最多的元素
func exercise2() {
	array := [5]int{1, 2, 2, 3, 2}
	countMap := make(map[int]int)
	for _, v := range array {
		countMap[v]++
	}
	maxCount := 0
	maxValue := 0
	for k, v := range countMap {
		if v > maxCount {
			maxCount = v
			maxValue = k
		}
	}
	fmt.Println("出现最多次数的元素:", maxValue)
	fmt.Println("countMap:", countMap)
}

// 小练习3：数组的冒泡排序

func exercise3() {
	array := [5]int{1, 4, 2, 3, 5}
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println("排序后的数组:", array)
}
