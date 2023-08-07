package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type ServerConfig struct {
	Port int `yaml:"port"`
}

type DBConfig struct {
	Postgres PostgresConfig `yaml:"postgres"`
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string
	Password string
	DbName   string
}

type Config struct {
	Env    string
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"db"`
}

func setField(obj interface{}, fieldName string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	fieldValue := structValue.FieldByName(fieldName)

	if !fieldValue.IsValid() {
		return fmt.Errorf("field with name %s not found", fieldName)
	}

	if !fieldValue.CanSet() {
		return fmt.Errorf("cannot set field with name %s", fieldName)
	}

	fieldValue.Set(reflect.ValueOf(value))
	return nil
}

func convertToFieldName(strVar string) string {
	if len(strVar) == 0 {
		return strVar
	}

	strArr := strings.Split(string(strVar), "_")
	var fieldName string
	for _, s := range strArr {
		firstChar := s[0]
		lowercased := strings.ToLower(s)
		chars := []byte(lowercased)
		chars[0] = firstChar
		fieldName += string(chars)
	}

	return fieldName
}

func loadEnvFile(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	parsed := strings.Split(string(data), "\n")
	var words []string
	for _, line := range parsed {
		if line == "" {
			continue
		}
		keys := strings.Split(string(line), "=")
		words = append(words, keys...)
	}

	for i, str := range words {
		if i%2 == 0 {
			words[i] = convertToFieldName(str)
		}
	}
	return words, nil
}

func GetConfig(path string) *Config {
	cfg := &Config{Env: os.Getenv("ENV")}
	data, err := loadEnvFile(path)

	if err != nil {
		panic(err)
	}

	for i, word := range data {
		if i%2 == 1 {
			continue
		}
		setField(cfg, word, data[i+1])
	}

	return cfg
}
