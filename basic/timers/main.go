package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个 2 秒后执行的 timer
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stopped := timer2.Stop()
	if stopped {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(time.Second * 2)
}
