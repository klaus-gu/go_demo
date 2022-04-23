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
		fmt.Println("im out ！！！")
	}()
	fmt.Println("START")
	//exit_by_os()
	exit_by_panic()
}
