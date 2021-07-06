package main

import (
	"context"

	"github.com/dink10/lessons/kafka/config"
	"github.com/dink10/lessons/kafka/consumer"
	"github.com/dink10/lessons/kafka/producer"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	cfg := config.LoadConfig()

	runConsumer(ctx, cfg)

	//runProducer(ctx, cfg)
}

func runConsumer(ctx context.Context, cfg *config.Config) {
	cons := consumer.NewConsumer([]string{cfg.Broker}, cfg.GroupID, cfg.Topic)

	for {
		logrus.Info("Start message reading")

		m, err := cons.FetchMessage(ctx)
		if err != nil {
			logrus.Errorf("error while getting kafka message: %v", err)
			continue
		}

		logrus.Infof(
			"[Partition %d] [Topic %s] Received key %s message: %s", m.Partition, m.Topic, m.Key, m.Value,
		)

		if err := cons.CommitMessages(ctx, m); err != nil {
			logrus.Error(err)
		}

		logrus.Info("Finish message reading")
	}
}

func runProducer(ctx context.Context, cfg *config.Config) {
	kafkaWriter := producer.NewProducer(ctx, cfg.Broker, cfg.Topic)

	// to create topics when auto.create.topics.enable='false'
	message := kafka.Message{
		Topic: cfg.Topic,
		Key:   []byte("Test-Key"),
		Value: []byte("Hello, World!"),
	}

	if err := kafkaWriter.Publish(ctx, message); err != nil {
		logrus.Fatal(err)
	}
}
