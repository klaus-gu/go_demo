package main

import "fmt"

func zeroVal(iVal int) {
	iVal = 0

}

/**
传入一个 int 指针引用，方法直接操作该引用对应的值
"*" 表示接受一个指针
"&" 表示使用一个指针
*/
func zeroPtr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("init val : ", i)
	zeroVal(i)
	fmt.Println("zeroVal: ", i)

	// &表示内存地址
	fmt.Println("&i:", &i)

	zeroPtr(&i)
	fmt.Println("zeroPtr: ", i)
	fmt.Println("pointer:", &i)

}
