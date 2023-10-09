package main

import (
	"container/list"
	"fmt"
)

func main() {
	ls := list.New()

	front := ls.PushFront(1)
	back := ls.PushBack(4)

	ls.InsertAfter(2, front)
	ls.InsertBefore(3, back)

	for e := ls.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
