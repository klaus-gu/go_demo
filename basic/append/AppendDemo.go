package main

import "fmt"

func main() {
	testAppendReturn()
}

func testAppendReturn() {
	var mySlice = make([]int, 5, 10)
	//sptr := &mySlice
	fmt.Printf("初始Slice内存地址：%p\n", &mySlice)

	ints := append(mySlice, 1)
	iptr := &ints
	fmt.Println("append之后的内存地址：", &iptr)
}
