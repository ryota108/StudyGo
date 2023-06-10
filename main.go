package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var m = map[string]string{
	"original1": "東京麺珍亭本舗",
	"original2": "モンスターズキッチン",
}

func main() {
	r := gin.Default()
	r.GET("/user/:id", handleUserRequest)
	r.GET("/restaurants/:id", restaurantHandle)
	r.GET("/restaurants", handleAllRestaurants)
	r.POST("/restaurants", addRestaurant)
	r.Run(":8080")
}

func handleUserRequest(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	dummyArray := []string{"Ryota", "Aya", "Sora"}
	dummyArray = append(dummyArray, "Tanaka")
	if num < 0 || num >= len(dummyArray) {
		c.JSON(http.StatusNotFound, gin.H{"error": "ID not found"})
		return
	}

	response := Response{ID: dummyArray[num]}
	c.JSON(http.StatusOK, response)
}

func restaurantHandle(c *gin.Context) {
	id := c.Param("id")
	response := Response{ID: id, Title: m[id]}
	c.JSON(http.StatusOK, response)
}

func handleAllRestaurants(c *gin.Context) {
	restaurants := make([]Response, 0, len(m))
	for id, title := range m {
		restaurants = append(restaurants, Response{ID: id, Title: title})
	}
	c.JSON(http.StatusOK, restaurants)
}

func addRestaurant(c *gin.Context) {
	var requestBody Response
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	m[requestBody.ID] = requestBody.Title
	c.JSON(http.StatusCreated, gin.H{"message": "Restaurant added successfully"})
}
