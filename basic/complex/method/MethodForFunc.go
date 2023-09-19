package method

import "fmt"

/*
*
自定义 type 类型，
*/
type my_func_type func(string)

func InArgs(funcType my_func_type) {
	str := "abc"
	funcType(str)
}

func OutArgs(str string) my_func_type {
	funct := func(string2 string) {
		fmt.Println(string2)
	}
	funct(str)
	return funct
}
