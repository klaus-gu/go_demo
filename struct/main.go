package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPersonConstruct(name string) *person {
	p := person{name: name}
	p.age = 2
	return &p
}

func main() {
	fmt.Println(person{
		name: "bob",
		age:  23,
	})

	s := newPersonConstruct("zs")
	s.age = 22
	fmt.Println(&s)
}
