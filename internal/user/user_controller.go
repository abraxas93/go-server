package user

import (
	"errors"
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
	var user *User
	var err error
	userID, err := strconv.Atoi(r.Params[":id"])
	// FIX: should throw invalid params error in case, when :id can't be converted to int
	if err != nil {
		json, _ := controller.GetJsonResponse(user, errors.New("InvalidUrlParams"))
		body = json
		w.WriteHeader(http.StatusNotAcceptable)
	}

	user, err = c.service.GetUser(userID)

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

	if user != nil {
		json, _ := controller.GetJsonResponse(user, nil)
		body = json
		w.WriteHeader(http.StatusOK)
	}
	// Write the JSON data to the response
	w.Write(body)
}
