package ka

import (
	"context"
	"github.com/segmentio/kafka-go"
)

var brokers = []string{"47.96.14.106:9092", "47.96.17.56:909", "116.62.164.174:9092"}

func NewProducer(topic string) (*kafka.Writer, *context.Context) {
	w := &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafka.Murmur2Balancer{},
		// 允许自动创建主题
		AllowAutoTopicCreation: true,
	}
	ctx := context.Background()
	return w, &ctx
}

func Send(key string, value string, producer *kafka.Writer, ctx context.Context) error {
	err := producer.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(value),
		})
	return err
}
