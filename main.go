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

func main() {
	r := gin.Default()
	r.GET("/user/:id", handleUserRequest)
	r.GET("/restaurants/:id", restaurantHandle)
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
	m := map[string]string{
		"J003451133": "仙台ホルモン焼肉酒場 ときわ亭 高田馬場店",
		"J001038613": "大衆酒場 ちばチャン 高田馬場店",
	}
	id := c.Param("id")
	response := Response{ID: id, Title: m[id]}
	c.JSON(http.StatusOK, response)
}
