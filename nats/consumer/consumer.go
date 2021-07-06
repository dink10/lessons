package consumer

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

type Consumer struct {
	*nats.Conn
}

func NewConsumer(url string) *Consumer {
	conn, err := nats.Connect(url)
	if err != nil {
		logrus.Fatal(err)
	}

	return &Consumer{
		Conn: conn,
	}
}
