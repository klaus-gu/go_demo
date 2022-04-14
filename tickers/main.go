package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个 每两秒执行一次的 Ticker
	ticker1 := time.NewTicker(2 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker1.C:
				fmt.Println("Tick at ", t)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	ticker1.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
