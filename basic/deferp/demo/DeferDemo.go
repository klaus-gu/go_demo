package demo

import "fmt"

func DeferOrder() {

	fmt.Println("准备测试 defer 执行顺序")

	defer func() {
		fmt.Println("第一个defer函数")
	}()

	defer func() {
		fmt.Println("第二个defer函数")
	}()

	defer func() {
		fmt.Println("第三个defer函数")
	}()
}

func DeferInDefer() {
	fmt.Println("准备测试 defer函数中执行一个 defer 函数")
	defer func() {
		fmt.Println("第一个defer函数")
		defer func() {
			fmt.Println("第二个defer函数内部的 defer")
		}()
	}()

	defer func() {
		fmt.Println("第二个defer函数")
		defer func() {
			fmt.Println("第二个defer函数内部的 defer 001")
		}()

		defer func() {
			fmt.Println("第二个defer函数内部的 defer 002")
		}()
	}()
}
