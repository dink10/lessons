package config

type Config struct {
	Host string `json:"host" env:"HOST"`
	Port int    `json:"port" env:"PORT"`
}
