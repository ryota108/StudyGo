package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user/{name}", handleUserRequest)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handleUserRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	response := Response{Message: "Hello, " + name + "!"}
	json.NewEncoder(w).Encode(response)
}

