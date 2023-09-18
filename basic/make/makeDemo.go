package main

import "fmt"

func main() {
	makeRtType()
}

func makeRtType() {
	var mySlice = make([]int, 5, 10)
	mySlice = append(mySlice, 1)
	var ptr = &mySlice
	fmt.Printf("mySlice类型：%T\n", mySlice)
	fmt.Println("mySlice内存地址：", &ptr)
	fmt.Printf("ptr是什么：%T\n", ptr)
	sliceB := &ptr
	//sliceB = append(sliceB, 2)
	fmt.Println("sliceB: ", *sliceB)
	fmt.Println("mySlice: ", &mySlice)

}
