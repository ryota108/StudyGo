package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Response struct {
	ID string `json:"id"`
	Title string `json:"title"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", handleUserRequest)
	r.HandleFunc("/restaurants/{id}", restaurantHandle)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handleUserRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	num, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// 配列のインデックスとして使用する
	dummyArray := []string{"Ryota", "Aya", "Sora"}
	dummyArray = append(dummyArray, "Tanaka")
	if num < 0 || num >= len(dummyArray) {
		http.Error(w, "ID not found", http.StatusNotFound)
		return
	}
	response := Response{ID: dummyArray[num]}
	json.NewEncoder(w).Encode(response)
}

func restaurantHandle(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	//dummy data
	m["J003451133"] = "仙台ホルモン焼肉酒場 ときわ亭 高田馬場店"
	m["J001038613"] = "大衆酒場 ちばチャン 高田馬場店"
	vars := mux.Vars(r)
	id := vars["id"]
	response := Response{ID:id,Title: m[id]}
	json.NewEncoder(w).Encode(response)
}
