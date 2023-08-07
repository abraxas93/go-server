package config

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Port int `yaml:"port"`
}

type DBConfig struct {
	Postgres PostgresConfig `yaml:"postgres"`
}

type PostgresConfig struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	DbUser     string
	DbPassword string
	DbName     string
}

type Config struct {
	Env    string
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"db"`
}

func loadEnvFile(filePath string) (map[string]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	parsed := strings.Split(string(data), "\n")
	keys := make(map[string]string, 0)
	for _, line := range parsed {
		if line == "" {
			continue
		}
		keyValue := strings.Split(string(line), "=")
		keys[keyValue[0]] = keyValue[1]

	}

	return keys, nil
}

func GetConfig(path string) *Config {
	cfg := &Config{Env: os.Getenv("ENV")}
	envKeys, err := loadEnvFile(path)

	if err != nil {
		panic(err)
	}

	filename := "config/local.yaml"
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = yaml.Unmarshal(data, cfg)

	if err != nil {
		fmt.Println(err)
	}

	cfg.DB.Postgres.DbUser = envKeys["POSTGRES_USER"]
	cfg.DB.Postgres.DbPassword = envKeys["POSTGRES_PASSWORD"]
	cfg.DB.Postgres.DbName = envKeys["POSTGRES_DB_NAME"]
	return cfg
}
