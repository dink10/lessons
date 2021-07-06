package producer

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

type Producer struct {
	*nats.Conn
}

func NewProducer(url string) *Producer {
	conn, err := nats.Connect(url)
	if err != nil {
		logrus.Fatal(err)
	}

	return &Producer{
		Conn: conn,
	}
}
