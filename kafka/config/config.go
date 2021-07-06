package config

import (
	"github.com/koding/multiconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Broker  string `default:"localhost:9092"`
	Topic   string `default:"test-kafka"`
	GroupID string `default:"consumer-group-1"`
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
