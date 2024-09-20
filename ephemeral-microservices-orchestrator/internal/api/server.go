package api

import (
	"fmt"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/create", createHandler)
	http.ListenAndServe(":8080", nil)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Microservice creation request received")
	// Logic to create microservice
}
