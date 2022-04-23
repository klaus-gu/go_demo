package main

import (
	"fmt"
	"sort"
)

func main() {
	stringArr := []string{"a", "c", "r", "s"}
	sort.Strings(stringArr)
	fmt.Println(stringArr)
	intArr := []int{2, 6, 3, 9}
	sort.Ints(intArr)
	fmt.Println(intArr)
}
