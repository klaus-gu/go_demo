package main

import "fmt"

type base struct {
	num int
}

func (b base) desc() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{num: 1},
		str:  "hahaha",
	}

	fmt.Printf("co={num: %v,str: %v}\n", co.num, co.str)
	fmt.Println("also num :", co.base.num)

	fmt.Println("desc:", co.desc())

	type describer interface {
		desc() string
	}

	/**
	有点类似Java的向下造型，父类的引用指向子类的对象
	*/
	var d describer = co
	fmt.Println("describer:", d.desc())

}
