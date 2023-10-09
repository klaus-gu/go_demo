package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	//wait_add_togther_done()
	wait_mutiAnd()
}

func wait_mutiAnd() {
	wg := sync.WaitGroup{}
	wg.Add(5)

	for w := 0; w < 5; w++ {
		fmt.Println("Finish: ", w)
		wg.Done()
	}
	wg.Wait()
	fmt.Println("结束程序")
}

func wait_add_togther_done() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	wg.Wait()
}
