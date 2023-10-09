package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//cancel()
	concurrentAdd()
}

func concurrentAdd() {
	count := 1
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 1; i <= 10000; i++ {
			count++
		}
		wg.Done()
	}()

	go func() {
		for i := 1; i <= 10000; i++ {
			count++
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(count)
}

func demo() {
	var ops uint64
	var wg sync.WaitGroup
	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go func() {
			for j := 1; j <= 1000; j++ {
				// 原子操作与非原子操作，二者计算时间差异，以及计算安全的差异
				atomic.AddUint64(&ops, 1)
				// ops++
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(ops)
}
