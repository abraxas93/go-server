package main

import (
	"fmt"
	"go-server/pkg/config"
	"go-server/pkg/logger"
	"os"

	"gopkg.in/yaml.v3"
)

type Cfg struct {
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"db"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type DBConfig struct {
	Postgres PostgresConfig `yaml:"postgres"`
}

type PostgresConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func main() {
	cfg := config.GetConfig(".env")
	logger.InitLogger(make(map[string]string))
	logger, _ := logger.GetLogger()
	logger.Info("%+v\n", cfg)
	// postgres.Connect()
	filename := "config/local.yaml"
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	c := &Cfg{}
	err = yaml.Unmarshal([]byte(data), c)

	if err != nil {
		fmt.Println(err)
	}
	logger.Info("%+v\n", c)

}
