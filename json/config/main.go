package main

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	DbHost     string
	DbPort     int
	DbName     string
	DbPassword string `xml:"pass" yaml:"password" json:"-"`
	DbUser     string `json:"-"`
}

func main() {
	cfg := Config{}

	// config.Read(&cfg) // yaml

	fmt.Println(json.Marshal(cfg))
}
