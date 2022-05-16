package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const RESPONSE_STRING = "欢迎来到德莱联盟！"

func logging(f func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		f(writer, request)
	}
}

func write_string(writer http.ResponseWriter, request *http.Request) {
	res, _ := json.Marshal(RESPONSE_STRING)
	writer.Write(res)
}

func write_string1(writer http.ResponseWriter, request *http.Request) {
	res, _ := json.Marshal("哈哈哈哈哈")
	writer.Write(res)
}

func write_struct_to_res(writer http.ResponseWriter, request *http.Request) {
	writer.Write(write_struct())
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/gorrila", logging(write_string))
	r.HandleFunc("/haha", logging(write_string1))
	r.HandleFunc("/getStruct", logging(write_struct_to_res))
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
