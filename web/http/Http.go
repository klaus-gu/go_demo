package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/decode", func(writer http.ResponseWriter, request *http.Request) {
		decode(request)
		fmt.Fprintf(writer, "Hello, you've requested: %s\n", request.URL.Path)
	})

	http.HandleFunc("/encode", func(writer http.ResponseWriter, request *http.Request) {
		encode(writer)
	})
	http.ListenAndServe(":80", nil)
}

/**
将请求json参数解析成 struct
*/
func decode(r *http.Request) {
	var person Person
	json.NewDecoder(r.Body).Decode(&person)
	fmt.Printf("请求用户：%s,年龄：%d", person.Name, person.Age)
}

/**
将 结构体转换成json
*/
func encode(w http.ResponseWriter) {
	peter := Person{
		Name: "peter",
		Age:  22,
	}
	json.NewEncoder(w).Encode(peter)
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
