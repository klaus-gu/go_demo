package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// 初始化
	r := ring.New(5)
	r.Value = 1

	// 赋值
	r.Next().Value = 2
	r = r.Next()
	r.Next().Value = 3
	r = r.Next()
	r.Next().Value = 4
	r = r.Next()
	r.Next().Value = 5

	// 遍历环形数组
	r.Do(func(a any) {
		fmt.Println(a)
	})

	s := ring.New(2)
	s.Value = 11
	s.Next().Value = 22

	// 将环形数组连接到另一个环形数组上
	r.Link(s)

	r.Do(func(a any) {
		fmt.Println(a)
	})

	// 从 r.next 开始解绑 11 个元素（这个例子中会把所有元素都删除）
	r.Unlink(11)
	r.Do(func(a any) {
		fmt.Println(a)
	})
}
