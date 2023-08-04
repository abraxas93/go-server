package main

import (
	"go-server/internal/logger"
	"log"
)

type Config struct {
	Port int
	Env  string
}

func main() {
	// Read the .env file
	// data, err := os.ReadFile(".env")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(strings.Split(string(data), "="))
	// cfg := Config{Port: 8080, Env: "dev"}
	// structValue := reflect.ValueOf(&cfg).Elem()
	// // structType := structValue.Type()
	// fmt.Println("Struct Type:", structValue)
	// fieldName := "Port"
	// fieldValue := reflect.ValueOf(cfg).FieldByName(fieldName)
	// if fieldValue.IsValid() {
	// 	fmt.Println(fieldName, "=", fieldValue.Interface())
	// } else {
	// 	fmt.Println(fieldName, "not found.")
	// }
	flags := log.LstdFlags | log.Lshortfile
	// infoLogger := log.New(os.Stdout, "", flags)
	// infoLogger.Printf("This is my string %q\n", "custom string")
	msgs := make(map[string]string)
	msgs["Info"] = "INFO: "
	msgs["Warn"] = "WARN: "
	msgs["Err"] = "ERROR: "
	l := logger.CreateLogger(msgs, flags)
	// fmt.Printf("%T\n", "sdfsg")

	l.Info("---> %T", "Hello World")
	l.Warn("---> %T", "Hello World")
	l.Err("---> %T", "Hello World")
	// l.Warn("Warning")
}
