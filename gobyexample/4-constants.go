package main

import (
	"fmt"
	"math"
)

const s string = "constant"

// 声明一个不被修改的常量值

func main() {
	fmt.Println(s)

	const n = 100
	const d = 200 / n
	fmt.Println(d)
	const u = 50000000
	const v = 3e20 / n
	fmt.Println(v, int64(v))
	fmt.Println(math.Sin(v))
}
