package complex

import "fmt"

type MyInterface interface {
	print(comment string) (result string)
}

type MyInterfaceImpl struct {
}

func (m MyInterfaceImpl) print(comment string) (result string) {
	fmt.Println("[ MyInterfaceImpl ] print: ", comment)
	return comment
}
