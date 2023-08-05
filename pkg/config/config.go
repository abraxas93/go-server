package config

import "os"

type Config struct {
	Env string
}

var cfg = Config{Env: os.Getenv("ENV")}

func GetConfig() Config {
	return cfg
}
