package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its s weekend")
	default:
		fmt.Println("Its a weekday")
	}

	fmt.Println(time.Now())

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Im a bool")
		case int:
			fmt.Println("Im a int")
		default:
			fmt.Printf("Dont know type %T", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("ehllo")
}
