package main

import (
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func main() {
	// cfg := config.GetConfig(".env")
	// log := logger.GetLogger()

	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	// Start the server
	http.ListenAndServe(":8080", mux)

}
