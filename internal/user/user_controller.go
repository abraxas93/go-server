package user

import (
	"encoding/json"
	"fmt"
	"go-server/pkg/utils/router"
	"net/http"
	"strconv"
	"strings"
)

type UserCtrl struct {
	service UserServiceIface
}

func NewUserCtrl(s UserServiceIface) *UserCtrl {
	return &UserCtrl{service: s}
}

func (c *UserCtrl) HelloHandler(w http.ResponseWriter, r *router.Request) {
	fmt.Println(r.Method)
	w.Write([]byte("Hello, world!"))
}

func (c *UserCtrl) AlbumHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	w.Write([]byte("Hello, world!"))
}

func (c *UserCtrl) CommentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	w.Write([]byte("Hello, world!"))
}

func (c *UserCtrl) GetUserHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL)
	fmt.Println(r.Method)
	parts := strings.Split(r.URL.Path, "/")
	fmt.Println(len(parts))
	userID, err := strconv.Atoi(parts[2])
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error while parsing", http.StatusBadRequest)
		return
	}
	fmt.Println(parts)
	// Marshal the User struct into JSON
	user, err := c.service.GetUser(userID)
	if err != nil {
		http.Error(w, "Something with postgres", http.StatusInternalServerError)
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
