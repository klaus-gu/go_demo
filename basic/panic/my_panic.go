package main

import (
	"errors"
	"fmt"
	"os"
)

func exit_by_os() {
	os.Exit(-1)
}

func exit_by_panic() {
	panic(errors.New("exit from panic"))
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panicWithDeferFunc()
}

func panicWithDeferFunc() {
	defer func() {
		fmt.Println("defer函数执行")
	}()
	exit_by_panic()
	fmt.Println("恢复执行")
}
