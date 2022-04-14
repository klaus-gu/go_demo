package main

import "fmt"

func main() {
	//create_channel()
	//close_channel()
	channel_demo()
}

func channel_demo() {
	myChannel := make(chan int, 5)
	myChannel <- 1
	myChannel <- 2
	myChannel <- 3
	myChannel <- 4
	myChannel <- 5
	close(myChannel)

	for elem := range myChannel {
		fmt.Println(elem)
	}
}

/**
关闭channel
*/
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

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job ", j)
	}
	close(jobs)
	// 重复关闭一个已经 close 的 channel 会出现 panic
	// close(jobs)
	fmt.Println("sent all jobs")
	<-done
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
