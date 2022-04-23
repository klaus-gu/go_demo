package main

import "fmt"

func main() {
	/**
	创建同一种类型的切片，
	创建空的切片需要使用 make
	*/
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	/**
	有点类似于 Java 的 List 的 add
	*/
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	/**
	没搞懂复制和直接赋值的区别。。。
	*/
	c := make([]string, len(s))
	// copy(c, s)
	// fmt.Println("copy:", c)
	c = s
	fmt.Println("copy:", c)

	// 从 2 复制到 (5-1) 左包含的概念
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	// 左包含的概念
	l = s[2:]
	fmt.Println("sl3:", l)

	// 区别于创建空的切片
	t := []string{"g", "h", "i"}
	fmt.Println("dcl：", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
