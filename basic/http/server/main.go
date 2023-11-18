package main

import (
	"fmt"
	"net/http"
)

func main() {
	myHandler := new(HelloWorldHandler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: myHandler,
	}
	http.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("foo")
	})
	server.ListenAndServe()
}

type HelloWorldHandler struct {
}

func (m *HelloWorldHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("处理请求")
	url := request.URL
	fmt.Println("请求路径：", url.Path)
}
