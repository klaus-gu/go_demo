package main

import (
	"fmt"
	"go_demo/basic/complex/method"
)

func main() {
	//methodForFuncIn()
	methodForFuncOut()
}

func methodForFuncIn() {
	funct := func(str string) {
		fmt.Println(str)
	}
	method.InArgs(funct)
}

func methodForFuncOut() {
	method.OutArgs("你好")
}
