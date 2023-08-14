package controller

import (
	"encoding/json"
)

type Response[T any] struct {
	Data T      `json:"data"`
	Err  string `json:"error"`
}

func GetJsonResponse[T any](data T, err error) ([]byte, error) {
	r := Response[any]{
		Data: data,
		Err:  err.Error(),
	}
	json, err := json.Marshal(r)
	return json, err
}
