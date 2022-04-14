package main

import "fmt"

func main() {
	create_channel()
}

func close_channel() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job :", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()
}

func create_channel() {
	// 创建一个叫做 messages 的 channel
	messages := make(chan string)
	// 将 ping 字符串 传入 messages channel
	go func() { messages <- "ping" }()
	// 输出 messages channel 的内容
	msg := <-messages
	fmt.Println(msg)

	// 创建一个 可以缓冲的 channel
	bufferedMessages := make(chan string, 2)
	bufferedMessages <- "hello"
	bufferedMessages <- "world"
	fmt.Println(<-bufferedMessages)
	fmt.Println(<-bufferedMessages)

	/**
	关闭 channel
	*/
	close(messages)
}
