package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	/*
		Range 在 arrays 和 slices 里面返回的是 index 下标和 value
	*/
	for index, value := range nums {
		sum += value
		fmt.Printf("Index: %d , Value: %d , Sum: %d .\n", index, value, sum)
	}

	mp := make(map[string]int)
	mp["one"] = 1
	mp["two"] = 2
	mp["three"] = 3
	/**
	在 Map 中，range 返回的是 key 和 value
	*/
	for key, value := range mp {
		fmt.Printf("key: %s ==>> value: %d.\n", key, value)
	}

}
