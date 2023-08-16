package controllers

import (
	"errors"
	"go-server/internal/user"
	"go-server/pkg/utils/controller"
	"go-server/pkg/utils/router"
	"net/http"
	"strconv"
)

type UserCtrl struct {
	service user.UserServiceIface
}

func NewUserCtrl(s user.UserServiceIface) *UserCtrl {
	return &UserCtrl{service: s}
}

func (c *UserCtrl) GetUserHandler(w http.ResponseWriter, r *router.Request) {
	// Set the appropriate Content-Type header
	w.Header().Set("Content-Type", "application/json")

	var body []byte
	var user *user.User
	var err error

	userID, err := strconv.Atoi(r.Params[":id"])

	if err != nil {
		json, _ := controller.GetJsonResponse(user, errors.New("InvalidUrlParams"))
		body = json
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write(body)
		return
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
