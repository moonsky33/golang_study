package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)
	// 打印map里的值，下面同理

	v1 := m["k1"]
	// 赋值v1为map里的k1
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	// 赋值v3为map里的k3，但是k3没有，所以v3为0
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	// 删除map里的k2
	fmt.Println("map:", m)

	clear(m)
	// 清空map里的值
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n==n2")
	}

}
