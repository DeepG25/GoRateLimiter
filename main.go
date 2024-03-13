package main

import (
	"log"
	"net/http"
	"ratelimiter/middleware"
	"ratelimiter/models"
	"time"

	"github.com/gin-gonic/gin"
)

func pingHandler(c *gin.Context) {

	message := "Health Ping triggered at " + time.Now().String()
	pingMessage := models.Response{Message: message}

	c.JSON(http.StatusOK, pingMessage)
}

func main() {
	router := gin.New()
	router.Use(middleware.RateLimiter(2, 4))

	router.GET("/ping", pingHandler)
	err := router.Run(":8080")
	if err != nil {
		log.Println("There was an error listening to port :8080", err)
	}
}
