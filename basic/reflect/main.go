package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.20
	fmt.Println("type: ", reflect.TypeOf(num))
	fmt.Println("value: ", reflect.ValueOf(num))
}
