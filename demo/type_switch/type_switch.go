package main

import (
	"fmt"
)

func main() {
	var t interface{}
	t = findType()
	switch t := t.(type) {
	case bool:
		fmt.Printf("这是一个布尔类型：%T", t)
	}
	fmt.Println(3 >> 2)
}

func findType() interface{} {
	return true
}
