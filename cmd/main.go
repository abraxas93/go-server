package main

import (
	"fmt"
	"go-server/pkg/logger"
	"log"
	"strings"
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
	str := "ENV"
	c := str[0]
	s := strings.ToLower("ENV")
	bytes := []byte(s)
	bytes[0] = c
	fmt.Println(string(bytes))
	flags := log.LstdFlags | log.Lshortfile

	msgs := make(map[string]string)
	msgs["Info"] = "INFO: "
	msgs["Warn"] = "WARN: "
	msgs["Err"] = "ERROR: "
	logger.InitLogger(msgs, flags)
	// fmt.Printf("%T\n", "sdfsg")
	logger, err := logger.GetLogger()
	if err != nil {
		fmt.Println(err)
	}
	logger.Info("---> %T", "Hello World")
	logger.Warn("---> %T", "Hello World")
	logger.Err("---> %T", "Hello World")
}
