package ka

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(brokers []string, topic string, group string) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: group,
	})
	consumer := new(Consumer)
	consumer.reader = reader
	return consumer
}

// ReadMessageByAutoCommitOffset 自动提交位移
func (consumer *Consumer) ReadMessageByAutoCommitOffset(ctx context.Context) (kafka.Message, error) {
	message, err := consumer.reader.ReadMessage(ctx)
	return message, err
}

// ReadMessageByExplicitCommitOffset 手动提交位移
func (consumer *Consumer) ReadMessageByExplicitCommitOffset(ctx context.Context) (kafka.Message, error) {
	message, err := consumer.reader.ReadMessage(ctx)
	return message, err
}

// CommitOffset 提交位移
func (consumer *Consumer) CommitOffset(message *kafka.Message, ctx *context.Context) {
	consumer.CommitOffset(message, ctx)
}
