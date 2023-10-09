package timeout

import (
	"context"
	"fmt"
	"time"
)

/*
*
设置指定的 超时时间之后，到时间会主动调用cancel函数
*/
func TimeOut() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	go func(ctx2 context.Context) {
		context.AfterFunc(ctx2, func() {
			fmt.Println("something done")
		})
		for {
			select {
			case <-ctx.Done():
				fmt.Println("i know ctx is done, and i will return")
				return
			default:
				fmt.Println("do something")
			}
			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("ctx is done")
	}
}
