package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/api", handleAPIRequest)
	http.HandleFunc("/test",testHandle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
// ok

func handleAPIRequest(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello, World!"}
	json.NewEncoder(w).Encode(response)
}

func testHandle(w http.ResponseWriter, r *http.Request){
	message := "Ok"
	response := Response{Message: message}
	json.NewEncoder(w).Encode(response)
}