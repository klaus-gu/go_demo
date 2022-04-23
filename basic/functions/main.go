package main

import "fmt"

/**
func 方法名(参数名 类型，参数名 类型) 返回值类型{}
*/
func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	plus := plus(1, 2)
	fmt.Printf("%d + %d = %d .\n", 1, 2, plus)
}
