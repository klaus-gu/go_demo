package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

const RESPONSE_STRING = "欢迎来到德莱联盟！"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/gorrila", func(writer http.ResponseWriter, request *http.Request) {
		res, _ := json.Marshal(RESPONSE_STRING)
		writer.Write(res)
	})
	r.HandleFunc("/haha", func(writer http.ResponseWriter, request *http.Request) {
		res, _ := json.Marshal("哈哈哈哈哈")
		writer.Write(res)
	})
	r.HandleFunc("/getStruct", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(write_struct())
	})
	http.ListenAndServe(":8888", r)
}

func write_struct() []byte {
	res := response2{
		Page:   0,
		Fruits: []string{"apple", "banana"},
	}
	byt, _ := json.Marshal(res)
	return byt
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}
