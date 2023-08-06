package main

import (
	"fmt"
	"go-server/pkg/config"
)

type Config struct {
	Port int
	Env  string
}

func main() {
	cfg := config.GetConfig(".env")
	fmt.Printf("%+v\n", cfg)
	// fmt.Println(convertToFieldName("POSTGRES_DB_USER"))
	// // Read the .env file
	// data, err := os.ReadFile(".env")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(data))
	// fmt.Println(config.GetConfig())
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
	// flags := log.LstdFlags | log.Lshortfile
	// // infoLogger := log.New(os.Stdout, "", flags)
	// // infoLogger.Printf("This is my string %q\n", "custom string")
	// msgs := make(map[string]string)
	// msgs["Info"] = "INFO: "
	// msgs["Warn"] = "WARN: "
	// msgs["Err"] = "ERROR: "
	// logger.InitLogger(msgs, flags)
	// // fmt.Printf("%T\n", "sdfsg")
	// logger, err := logger.GetLogger()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// logger.Info("---> %T", "Hello World")
	// logger.Warn("---> %T", "Hello World")
	// logger.Err("---> %T", "Hello World")
}
