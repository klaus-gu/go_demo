package main

import "fmt"

func main() {
	// 值类型参数，实参的值未改变
	mySalary := 80000
	fmt.Printf("变量 mySalary 的内存地址为：%p\n", &mySalary)

	modifySalary(mySalary)
	fmt.Println(mySalary)

	// 指针类型参数，实参的值被改变
	modifySalary2(&mySalary)
	fmt.Println(mySalary)
}

func modifySalary(salary int) {
	fmt.Printf("参数变量的内存地址为：%p\n", &salary)
	salary = 100000
}

func modifySalary2(salary *int) {
	fmt.Printf("参数变量的内存地址为：%p\n", salary)
	*salary = 100000
}
