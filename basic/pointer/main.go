package main

import "fmt"

func zeroVal(iVal int) {
	iVal = 0

}

/*
*
传入一个 int 指针引用，方法直接操作该引用对应的值
"*" 表示接受一个指针
"&" 表示使用一个指针
*/
func zeroPtr(iptr *int) {
	*iptr = 0
}

func main() {
	struct_display()
}

type Person struct {
	name string
}

func (p Person) getName() string {
	return p.name
}

/**
暂定认为：

&：是获取当前类型指针的地址
*：是获取指针地址所存储的值

由于 go 是会自动获取地址的值的，所以就算不用 * 去取值，也会主动被取出来
*/

func struct_display() {

	p := &Person{
		name: "nike",
	}
	// 将 p 的指针地址 赋值给 b
	b := p

	fmt.Println(p)
	fmt.Println(b)
}

func base_display() {
	//i := 1
	//fmt.Println("init val : ", i)
	//zeroVal(i)
	//fmt.Println("zeroVal: ", i)
	//
	//// &表示内存地址
	//fmt.Println("&i:", &i)
	//
	//zeroPtr(&i)
	//fmt.Println("zeroPtr: ", i)
	//fmt.Println("pointer:", &i)
}
