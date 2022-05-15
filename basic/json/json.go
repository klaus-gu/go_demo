package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	/**
	json.Marshal: 将某一种类型的数据结构转换成 byte[]
	在通过 string() 函数进行转换
	*/
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(3.14)
	fmt.Println(string(fltB))

	slcD := []string{"apple", "banana", "peach"}
	slcj, _ := json.Marshal(slcD)
	fmt.Println(string(slcj))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}
