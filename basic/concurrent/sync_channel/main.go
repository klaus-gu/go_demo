package main

import (
	"fmt"
	"time"
)

func worker(done chan<- bool) {
	fmt.Println("working ... ")
	time.Sleep(5 * 1000 * time.Millisecond)
	fmt.Println("done")
	// 执行完毕之后通知 channel 我已经结束了
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}
