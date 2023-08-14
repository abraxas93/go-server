package user

import (
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
	// Set the appropriate Content-Type header
	w.Header().Set("Content-Type", "application/json")
	var body []byte
	userID, err := strconv.Atoi(r.Params[":id"])

	if err != nil {
		http.Error(w, "Error while parsing", http.StatusBadRequest)
		return
	}

	// Marshal the User struct into JSON
	user, err := c.service.GetUser(userID)

	if user == nil {
		json, _ := controller.GetJsonResponse(user, errors.New("UserNotFound"))
		body = json
		w.WriteHeader(http.StatusBadRequest)
	}

	if err != nil {
		json, _ := controller.GetJsonResponse(user, errors.New("InternalServerError"))
		body = json
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println(user)
	if user != nil {
		json, err := controller.GetJsonResponse(user, nil)
		fmt.Println(string(json))
		fmt.Println(err)
		body = json
		w.WriteHeader(http.StatusOK)
	}
	// Write the JSON data to the response
	w.Write(body)
}
