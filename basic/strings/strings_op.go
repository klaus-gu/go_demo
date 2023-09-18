package main

import (
	"fmt"
	"strconv"
)

func main() {
	string_operate()
}

func string_operate() {
	a := 1
	var b string
	b = strconv.Itoa(a) + "你好"
	fmt.Println("数字：", a)
	fmt.Println("字符串拼接：", b)
}
