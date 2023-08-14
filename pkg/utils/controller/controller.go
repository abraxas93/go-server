package controller

import (
	"encoding/json"
	"fmt"
)

type Response[T any] struct {
	Data T     `json:"data"`
	Err  error `json:"error"`
}

func GetJsonResponse[T any](data T, err error) ([]byte, error) {
	fmt.Println(&data)
	fmt.Println(data)
	r := Response[any]{
		Data: data,
		Err:  err,
	}
	json, err := json.Marshal(r)
	return json, err
}
