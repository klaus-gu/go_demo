package main

import (
	"fmt"
	"sync"
)

/*
*
once 中定义的方法只会执行一次
*/
func main() {
	once := sync.Once{}
	once.Do(func() {
		fmt.Println("shuchu")
	})
	once.Do(func() {
		fmt.Println("shuchu")
	})
}
