package main

import "go_demo/middleware/kafka/ka"

func main() {
	//defer ka.MyConn.Close()
	//ka.CreateTopic("go-topic2")
	//time.Sleep(5 * time.Second)
	ka.ListTopic()
	//producer, ctx := ka.NewProducer("go-topic3")

}
