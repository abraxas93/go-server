package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-server/pkg/utils/controller"
	"go-server/pkg/utils/router"
	"net/http"
	"strconv"
)

type UserCtrl struct {
	service UserServiceIface
}

func NewUserCtrl(s UserServiceIface) *UserCtrl {
	return &UserCtrl{service: s}
}

func (c *UserCtrl) GetUserHandler(w http.ResponseWriter, r *router.Request) {

	userID, err := strconv.Atoi(r.Params[":id"])

	if err != nil {
		http.Error(w, "Error while parsing", http.StatusBadRequest)
		return
	}

	// Marshal the User struct into JSON
	user, err := c.service.GetUser(userID)

	if user == nil {
		json, _ := controller.GetJsonResponse(user, errors.New("UserNotFound"))
		fmt.Println(string(json))
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
		return
	}

	if err != nil {
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	// Set the appropriate Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON data to the response
	w.Write(jsonData)
}
