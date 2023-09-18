package main

import "fmt"

func main() {
	//user()
	//newInt()
	userByConstruct()
}

func userByConstruct() {
	var user = NewUserConstructor()
	fmt.Println("原User内存地址：", &user)
	fmt.Println("原User：", *user)
	user.setName("张三")
	fmt.Println("getName: ", user.getName())
}

func newInt() {
	var a = new(int)
	fmt.Println("内存地址：", &a)

}

/*
*
说明：new 创建返回的是指针类型的数据
*/
func user() {
	var user = NewUserConstructor()
	fmt.Println("内存地址：", &user)
	fmt.Println("取 new 的指针的值：", *user)
	fmt.Printf("type:%T\n", user)
	fmt.Printf("name:%s,age:%d\n", user.name, user.age)
	user.name = "zhangsan"
	user.age = 11
	fmt.Printf("name:%s,age:%d", user.name, user.age)
}
