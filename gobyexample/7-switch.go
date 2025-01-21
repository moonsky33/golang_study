package main

import "fmt"

// switch语句是不同条件执行不同代码，并且更加简洁易读，类似于多个if-else语句

// func main() {
// 	i := 2
// 	fmt.Println("Write", i, "as")
// 	switch i {
// 	case 1:
// 		fmt.Println("one")
// 	case 2:
// 		fmt.Println("two")
// 	case 3:
// 		fmt.Println("three")
// 	}
// 	switch time.Now().Weekday() {
// 	case time.Saturday, time.Sunday:
// 		fmt.Println("It's the weekend")
// 	default:
// 		fmt.Println("It's a weekday")
// 	}
// 	t := time.Now()
// 	switch {
// 	case t.Hour() < 12:
// 		fmt.Println("It's before noon")
// 	default:
// 		fmt.Println("It's after noon")
// 	}
// 	whatAmI := func(i interface{}) {
// 		switch t := i.(type) {
// 		case bool:
// 			fmt.Println("I'm a bool")
// 		case int:
// 			fmt.Println("I'm an int")
// 		default:
// 			fmt.Printf("Don't know type %T\n", t)
// 		}
// 	}
// 	whatAmI(true)
// 	whatAmI(1)
// 	whatAmI("hey")
// }

// 小练习1：实现根据用户输入的操作符和两个操作数进行计算的程序
// 编写一个go，程序接收三个输入：+、-、*、/、%和两个操作数，然后根据输入的操作符进行相应的计算

func calculate(operator string, a, b int) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	case "%":
		return a % b
	default:
		return 0
	}
}
func main() {
	var operator string
	var a, b int
	fmt.Print("请输入操作符 (+, -, *, /, %): ")
	fmt.Scan(&operator)
	fmt.Print("请输入第一个操作数: ")
	fmt.Scan(&a)
	fmt.Print("请输入第二个操作数: ")
	fmt.Scan(&b)
	result := calculate(operator, a, b)
	fmt.Printf("%d %s %d = %d\n", a, operator, b, result)
}
