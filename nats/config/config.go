package config

import (
	"github.com/koding/multiconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	NatsURL string `default:"nats://localhost:4222"`
	Subject string `default:"test-nats"`
}

func LoadConfig() *Config {
	var cfg Config

	m := multiconfig.New()
	err := m.Load(&cfg)
	if err != nil {
		logrus.Fatal(err)
	}

	return &cfg
}
