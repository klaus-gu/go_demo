package main

import (
	"context"
	"fmt"
	"go_demo/middleware/kafka/ka"
)

func main() {
	//defer ka.MyConn.Close()
	//ka.CreateTopic("go-topic2")
	//time.Sleep(5 * time.Second)
	//ka.ListTopic()

	producer := ka.NewProducerAssignedTopic("localhost:9092", "go-topic2")
	//defer func() { producer.Close() }()

	for i := 10; i > 0; i-- {
		producer.SendMessage("hello", "world")
	}
	consumer := ka.NewConsumer([]string{"localhost:9092"}, "go-topic2", "go-topic2-group1")
	for {
		offset, err := consumer.ReadMessageByAutoCommitOffset(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println("主题：", offset.Topic, "== 所属分区：", offset.Partition, " == key：", string(offset.Key), " == value: ", string(offset.Value))
	}

}
