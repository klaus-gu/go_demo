package cancel

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func Test() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			fmt.Println(i)
			context.WithCancel(ctx)
			time.Sleep(5 * time.Second)
			wg.Done()
		}()
	}

	fmt.Println("当前活跃goroutine：", runtime.NumGoroutine())
	time.Sleep(2 * time.Second)

	cancelFunc()
	fmt.Println("调用 cancelFunc ...")
	fmt.Println("当前活跃goroutine：", runtime.NumGoroutine())
	wg.Wait()
}
