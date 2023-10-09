package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	// 遍历通道内容
	for j := range jobs {
		fmt.Println("worker ", id, "started job ", j)
		time.Sleep(time.Second)
		fmt.Println("worker ", id, "finished job", j)
		// 将数据写入
		results <- j
	}
}

func main() {
	const num = 5
	jobs := make(chan int, num)
	results := make(chan int, num)
	/**
	创建三个 worker
	*/
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	/**
	往 channel 中写入
	*/
	for j := 1; j <= num; j++ {
		jobs <- j
	}
	close(jobs)
	/**
	从 channel 中读取
	*/
	for a := 1; a <= num; a++ {
		<-results
	}
}
