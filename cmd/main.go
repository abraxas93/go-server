package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read the .env file
	data, err := os.ReadFile(".env")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(strings.Split(string(data), "="))

}
