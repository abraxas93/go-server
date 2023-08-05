package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Config struct {
	Env string
}

var cfg = Config{Env: os.Getenv("ENV")}

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
	c := strVar[0]
	l := strings.ToLower(strVar)
	bytes := []byte(l)
	bytes[0] = c
	return string(bytes)
}

func GetConfig() Config {
	fieldValue := reflect.ValueOf(cfg)
	return cfg
}
