package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	//create_errorString()
	//create_errorWrap()
	//create_unwrap()
	errors_is()
}

/**
判断这个 error 是否包含某一个特定的 error
*/
func errors_is() {
	err := fmt.Errorf("wrap err:%w", os.ErrPermission)
	if res := errors.Is(err, os.ErrPermission); res {
		fmt.Println("Yes is os err")
	}
}

/**
fmt.Errorf("***",error)：可以将error或者wrapError再一次包装，层层包装
errors.Unwrap(error):将error再次解开一次包装
*/
func create_unwrap() {
	err := fmt.Errorf("write file error:%w", os.ErrPermission)
	if errors.Unwrap(err) == os.ErrPermission {
		fmt.Println("permission denied")
	}

	my_unwrap_err := fmt.Errorf("again:%w", err)
	if err == errors.Unwrap(my_unwrap_err) {
		fmt.Println("yes")

	}
}

func create_errorString() {
	err := errors.New("this is a demo error.")
	basicErr := fmt.Errorf("some context: %v", err)

	if _, ok := basicErr.(interface{ Unwrap() error }); !ok {
		fmt.Print("basicErr is a errorString")
	}
}

func create_errorWrap() {
	err := errors.New("this is a demo error .")
	// fmt.Errorf 给 errorstring 套上一层皮
	basicErr := fmt.Errorf("some context:%w", err)
	/**
	语法：检查某一个 type 是否实现某一个接口
	type.(被实现的接口)
	示例：basicErr.(interface{ Unwrap() error })
	*/
	if _, ok := basicErr.(interface{ Unwrap() error }); ok {
		fmt.Print("basicError is a wrapError")
	}
	// basicErr.Error() 就是蜕去皮之后的原本的 errorstring
	fmt.Println(basicErr.Error())
}
