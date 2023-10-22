package main

import (
	"bytes"
	"fmt"
	"io"
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

	bt := []byte{}
	resp.Body.Read(bt)

	io.Copy(os.Stdout, bytes.NewReader(bt))

	resp.Body.Close()
}

func http_post() {

}
