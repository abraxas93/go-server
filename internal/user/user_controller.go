package user

import (
	"encoding/json"
	"net/http"
)

type UserCtrl struct {
	service UserServiceInterface
}

func (c *UserCtrl) HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func (c *UserCtrl) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Marshal the User struct into JSON
	user, err := c.service.GetUser(2)
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
