package consumer

import "github.com/segmentio/kafka-go"

type Consumer struct {
	*kafka.Reader
}

func NewConsumer(brokers []string, groupID string, topic string) *Consumer {
	config := kafka.ReaderConfig{
		Brokers: brokers,
		GroupID: groupID,
		Topic:   topic,
	}

	reader := kafka.NewReader(config)

	return &Consumer{Reader: reader}
}
