package user

import "net/http"

type UserCtrl struct {
	service *UserServiceInterface
}

func (c *UserCtrl) HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
