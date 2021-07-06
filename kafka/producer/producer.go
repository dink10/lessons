package producer

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type Producer struct {
	*kafka.Conn
}

func NewProducer(ctx context.Context, brokers string, topic string) *Producer {
	conn, err := kafka.DialLeader(ctx, "tcp", brokers, topic, 0)
	if err != nil {
		logrus.Fatal(err)
	}

	return &Producer{
		Conn: conn,
	}
}

func (p *Producer) Publish(ctx context.Context, message kafka.Message) error {
	_, err := p.WriteMessages(message)
	if err != nil {
		return err
	}

	return nil
}
