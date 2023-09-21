package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

// ==========================================   长方形   =========================================
type rect struct {
	width, height float64
}

/*
*
实现接口的函数
*/
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// ==========================================   圆   =========================================
type circle struct {
	radius float64
}

/*
*
实现接口的函数
*/
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// =================================  统一使用 ：类似模版方法的client类 ====================================
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
