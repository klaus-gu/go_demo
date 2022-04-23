package main

import "fmt"

func vals() (int, int) {
	return 3, 4
}

func main() {
	a, b := vals()
	fmt.Println(a, b)
}
