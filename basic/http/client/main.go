package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http_get()
}

func http_get() {
	resp, err := http.Get("http://127.0.0.1:8001/hello/base?name=%E5%88%98%E8%89%B3")
	if err != nil {
		fmt.Errorf(err.Error())
	}

	buff := new(bytes.Buffer)
	buff.Grow(1024)
	fmt.Println(buff.Cap())
	body := resp.Body
	//allBytes, err := io.ReadAll(body)
	//io.Copy(os.Stdout, bytes.NewReader(allBytes))
	for {
		fmt.Println("buff len: ", buff.Len())
		bt := make([]byte, 20)
		n, err := body.Read(bt)
		fmt.Println("write == ", n)
		if err != nil {
			fmt.Println("Has read all.")
			break
		} else {
			buff.Write(bt)
		}
	}
	fmt.Println("final len: ", buff.Len())
	buff.WriteTo(os.Stdout)
	resp.Body.Close()
}

func http_post() {

}
