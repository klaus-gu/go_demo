package main

import "fmt"

func main() {
	//slices()
	append_test()
}

func append_test() {
	var sliceA = make([]int, 5)
	for i := 1; i <= 5; i++ {
		sliceA[i-1] = i
	}
	fmt.Printf("初始化内存地址：%x", &sliceA)

	sliceA = append(sliceA, 9)
	fmt.Printf("append 之后的内存地址：%x", &sliceA)
}

func slices() {
	var sliceA = make([]string, 10)
	fmt.Println("初始大小：", len(sliceA))

	for i := 0; i < 10; i++ {
		sliceA[i] = "元素" + string(i)
	}
	var childSlice = sliceA[:6]
	fmt.Printf("子slice: %v\n", childSlice)
	for a, b := range childSlice {
		fmt.Printf("下标：%d，值：%s\n", a, b)
	}
}

func demo() {
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
