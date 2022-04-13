package main

import "fmt"

/**
函数是特殊的方法，函数前面的指针，指定了这个函数只能在某特定类型的指针后使用
*/
type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{
		width:  10,
		height: 5,
	}
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())
	fmt.Println(&r)

	rp := &r
	rp.width = 10000
	fmt.Println("area:", rp.area())
	fmt.Println("perim: ", rp.perim())
	fmt.Println(rp.width)

	fmt.Println(r.width)

	// 使用 * 赋值给一个新的结构，*XX表示引用，&XX表示获取XX的内存地址是一个取址符
	cpr := *rp
	cpr.width = 1002
	fmt.Println("cprw:", cpr.width)
	fmt.Println("cpw:", rp.width)
}
