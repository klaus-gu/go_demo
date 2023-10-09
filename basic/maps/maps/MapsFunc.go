package maps

import (
	"fmt"
	"maps"
)

var CommonMap = map[string]int{
	"zhangsan": 12,
	"lisi":     14,
	"wangwu":   33,
	"zhaoliu":  44,
}

func DeleteFunc() {
	maps.DeleteFunc(CommonMap, func(s string, i int) bool {
		if s == "wangwu" {
			return true
		} else {
			return false
		}
	})

	for k := range CommonMap {
		fmt.Printf("key: %s, value: %d\n", k, CommonMap[k])
	}

}

func CloneFunc() {

}

func CopyFunc() {

}

func EqualFunc() {

}

func EqualFuncFunc() {

}
