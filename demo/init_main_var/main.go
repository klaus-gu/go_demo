package main

import "fmt"

var tag = varFunc()

var varFunc = func() string {
	fmt.Println("变量的初始化")
	return "change by varFunc"
}

func init() {
	fmt.Println("init...")
}

func main() {
	fmt.Println("main...")
}
