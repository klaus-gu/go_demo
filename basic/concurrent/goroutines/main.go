package main

import (
	"fmt"
	"go_demo/basic/concurrent/goroutines/communication"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	communication.Communicate()

	//f("direct")
	//
	//go f("goroutine")
	//
	//go func(msg string) { fmt.Println(msg) }("going")
	//
	//time.Sleep(time.Second)
	//
	//fmt.Println("done")
}
