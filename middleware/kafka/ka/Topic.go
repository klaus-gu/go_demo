package ka

import (
	"fmt"
	"github.com/segmentio/kafka-go"
)

func CreateTopic(topic string) {
	kafkaTopics := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}
	MyConn.CreateTopics(kafkaTopics...)
}

func ListTopic() {
	partitions, err := MyConn.ReadPartitions()
	if err != nil {
		fmt.Println(err.Error())
	}
	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k := range m {
		fmt.Println(k)
	}
}
