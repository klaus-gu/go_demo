package communication

import (
	"fmt"
	"time"
)

func Communicate() {

	myChannel := make(chan string)

	go func() {
		myChannel <- "hello"
	}()

	go func() {
		msg := <-myChannel
		fmt.Println("收到消息=", msg)
	}()

	time.After(time.Second * 1)
}
