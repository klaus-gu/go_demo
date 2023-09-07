package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
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

func test() {

}
