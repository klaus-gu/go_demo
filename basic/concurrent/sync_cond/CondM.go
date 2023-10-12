package main

import (
	"fmt"
	"sync"
	"time"
)

var done = false

func read(cond *sync.Cond) {
	cond.L.Lock()
	/**
	  if语句只会对共享资源的状态检查一次，而for语句却可以做多次检查，直到这个状态改变为止。
	  那为什么要做多次检查呢？这主要是为了保险起见。如果一个 goroutine 因收到通知而被唤醒，
	  但却发现共享资源的状态，依然不符合它的要求，那么就应该再次调用条件变量的Wait方法，
	  并继续等待下次通知的到来。
	*/
	for !done {
		fmt.Println("开始阻塞")
		cond.Wait()
	}
	fmt.Println("恢复执行")
	cond.L.Unlock()
}

func write(cond *sync.Cond) {
	fmt.Println("开始写入 ...")
	time.Sleep(time.Second)
	cond.L.Lock()
	done = true
	cond.L.Unlock()
	fmt.Println("写入完毕！")
	cond.Broadcast()
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	go read(cond)
	go read(cond)
	go write(cond)
	time.Sleep(5 * time.Second)
}
