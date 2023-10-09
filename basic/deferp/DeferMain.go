package main

import (
	"fmt"
	"go_demo/basic/deferp/demo"
)

/*
*

 1. defer 函数在同一层代码中是以 LIFO 的方式执行的，即根据定义 defer 的顺序倒叙执行
 2. defer 函数内部的 defer 函数执行顺序遵循上一条
*/
func main() {
	fmt.Println("======================Defer Order======================")
	demo.DeferOrder()

	fmt.Println("======================Defer In Defer======================")
	demo.DeferInDefer()
}
