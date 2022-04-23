package main

import "fmt"

/**
闭包，有点类似于Java的匿名函数
*/
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	netInt := intSeq()
	fmt.Println(netInt())
	fmt.Println(netInt())
	fmt.Println(netInt())
}
