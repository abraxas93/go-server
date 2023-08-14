package controller

import (
	"encoding/json"
)

type Response[T any] struct {
	Data T       `json:"data"`
	Err  *string `json:"error"`
}

func GetJsonResponse[T any](data T, err error) ([]byte, error) {
	errStr := ""
	if err != nil {
		errStr = err.Error()
	}

	var errPtr *string
	if errStr != "" {
		errPtr = &errStr
	}

	r := Response[any]{
		Data: data,
		Err:  errPtr,
	}

	json, err := json.Marshal(r)
	return json, err
}
