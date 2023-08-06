package main

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Port int
	Env  string
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

func main() {
	data, err := loadEnvFile(".env")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
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
