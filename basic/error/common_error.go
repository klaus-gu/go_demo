package main

import (
	"errors"
	"fmt"
)

var LessThanTwoErroe = errors.New("n should be not less than 2")
var LargerThanHundredError = errors.New("n should be not larger than 100")

func main() {
	//my_error()
	fail_fast_by_err(101)
}

func fail_fast_by_err(num int64) error {
	if num < 2 {
		return LessThanTwoErroe
	} else if num > 100 {
		return LargerThanHundredError
	}
	return nil
}

func fail_fast(num int64) (int64, error) {
	if num < 0 || num > 100 {
		return 0, errors.New("")
	}
	return num * 2, nil
}

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("cant work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg: arg, prob: "cant work with it"}
	}
	return arg + 3, nil
}

func my_error() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
