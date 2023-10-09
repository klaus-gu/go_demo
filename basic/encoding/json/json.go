package main

import (
	"encoding/json"
	"os"
)

/*
*
结构体中的 参数可以通过 `json:"name"` 的格式进行自定义格式化之后的json key
同时，通过在格式中声明 omitempty 关键字可以使得在序列化的时候忽略 false 0 nil 的字段
*/
func main() {
	type User struct {
		Name string `json:"myname"`
		Age  int    `json:"myage"`
	}

	zhangsan := User{"zhangsan", 12}
	b, _ := json.Marshal(zhangsan)
	os.Stdout.Write(b)

	// 设置当 Kinds 类型字段为空字符串的时候不在序列化的时候输出
	type Animal struct {
		Name    string `json:"name"`
		Kinds   string `json:"kinds,omitempty"`
		HasFoot bool   `json:"hasFoot"`
	}

	dog := Animal{
		"dog", "", true,
	}

	b1, _ := json.Marshal(dog)
	os.Stdout.Write(b1)
}
