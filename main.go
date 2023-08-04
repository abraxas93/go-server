package main

import (
	"fmt"
	"os"
)

func main() {
	// Read the .env file
	data, err := os.ReadFile(".env")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

}
