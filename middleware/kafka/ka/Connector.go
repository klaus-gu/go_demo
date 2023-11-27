package ka

import "github.com/segmentio/kafka-go"

var MyConn = connect()

func connect() *kafka.Conn {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	return conn
}