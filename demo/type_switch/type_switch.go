package main

import (
	"fmt"
	"go/types"
)

func main() {
	var t interface{}
	t = findType()
	switch t := t.(type) {
	case bool:
		fmt.Printf("这是一个布尔类型：%T", t)
	}
	fmt.Println()
}

func findType() interface{} {
	return types.Bool
}
