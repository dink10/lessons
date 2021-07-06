package main

import (
	"github.com/dink10/lessons/nats/config"
	"github.com/dink10/lessons/nats/producer"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.LoadConfig()

	pr := producer.NewProducer(cfg.NatsURL)

	if err := pr.Publish(cfg.Subject, []byte("Hello, World!")); err != nil {
		logrus.Fatal(err)
	}

	_, err := pr.Subscribe(cfg.Subject, func(msg *nats.Msg) {
		logrus.Infof("Received message: %s", string(msg.Data))
		if err := pr.Publish(msg.Reply, []byte("Done")); err != nil {
			logrus.Fatal(err)
		}
	})
	if err != nil {
		logrus.Fatal(err)
	}

	select {}
}
