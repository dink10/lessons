package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/zn11ch/SimpleWebApp/config"
	"github.com/zn11ch/SimpleWebApp/internal/app/apiserver"
)

func main() {

	// todo read config

	var cfg config.Config

	f, _ := os.Open("config.jon")
	configData, _ := io.ReadAll(f)

	_ = json.Unmarshal(configData, &cfg)

	api := apiserver.New(&cfg)
	api.Start()
}
