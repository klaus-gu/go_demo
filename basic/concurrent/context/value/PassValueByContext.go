package value

import (
	"context"
	"fmt"
	"sync"
)

func PassValue() {
	ctx := context.WithValue(context.Background(), "go", "goroutine-01")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(ctx1 context.Context) {
		fmt.Println(ctx1.Value("go"))
		wg.Done()
	}(ctx)
	wg.Wait()
}
