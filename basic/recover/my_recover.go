package main

import (
	"errors"
	"fmt"
)

func main() {
	recover_from_panic()
}

func recover_from_panic() {
	fmt.Println("STRAT...")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("RECOVERED FROM :", err)
		}
	}()

	panic(errors.New("something wrong!"))
}
