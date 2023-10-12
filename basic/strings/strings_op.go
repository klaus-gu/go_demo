package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//string_operate()
	//clone()
	r := []rune("你好")
	fmt.Println(len(r))
}

func clone() {
	s := "abc"
	s2 := strings.Clone(s)
	fmt.Println(s == s2)
	fmt.Println(&s)
	fmt.Println(&s2)
}

func string_operate() {
	a := 1
	var b string
	b = strconv.Itoa(a) + "你好"
	fmt.Println("数字：", a)
	fmt.Println("字符串拼接：", b)
}
