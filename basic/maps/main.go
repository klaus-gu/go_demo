package main

import (
	"fmt"
	"go_demo/basic/maps/maps"
)

func main() {
	//demo()
	maps.DeleteFunc()
}

func demo() {
	/**
	创建空的 map 类似 slices 使用 make 创建
	*/
	m := make(map[string]int)
	// 赋值
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map: ", m)

	// 根据 key 获取 value
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	// 删除
	delete(m, "k2")
	fmt.Println("map:", m)

	// 取值的时候有两个返回值：（值本身）和（该 key 是否存在于这个 map）
	_, prs := m["k2"]
	fmt.Println("prs", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
