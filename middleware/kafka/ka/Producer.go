package ka

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

/*
*
创建一个指定主题的生产者
*/
func NewProducerAssignedTopic(brokers string, topic string) *Producer {
	p := new(Producer)
	writer := &kafka.Writer{
		Addr:                   kafka.TCP(brokers),
		Topic:                  topic,
		AllowAutoTopicCreation: true,
		Compression:            kafka.Lz4,
		Balancer:               &kafka.Murmur2Balancer{},
	}
	p.writer = writer
	return p
}

/*
*
创建一个不指定主题的生产者
*/
func NewProducer(brokers string) *Producer {
	p := new(Producer)
	writer := &kafka.Writer{
		Addr:                   kafka.TCP(brokers),
		AllowAutoTopicCreation: true,
		Compression:            kafka.Lz4,
		Balancer:               &kafka.Murmur2Balancer{},
	}
	p.writer = writer
	return p
}

/*
*
发送消息
*/
func (producer *Producer) SendMessage(key string, value string) error {
	ctx := context.Background()
	err := producer.writer.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(value),
		})
	return err
}

/*
*
关闭生产者
*/
func (p *Producer) Close() error {
	return p.Close()
}

// var brokers = []string{"47.96.14.106:9092", "47.96.17.56:9092", "116.62.164.174:9092"}
var brokers = []string{"localhost:9092"}
