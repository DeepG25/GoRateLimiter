package middleware

import (
	"log"
	"net/http"
	"ratelimiter/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GetRateLimiter(inflowRate int, bucketSize int) models.RateLimiter {
	ratelimiter := &models.TokenBucketRateLimiter{
		InflowRate:             inflowRate,
		BucketSize:             bucketSize,
		TokenCount:             bucketSize,
		LastRefillTime:         time.Now().UTC(),
		RefreshWindowInSeconds: 10,
	}

	log.Println("Rate limiter initalized")
	return ratelimiter
}

func RateLimiter(inflowRate int, bucketSize int) gin.HandlerFunc {

	rateLimiter := GetRateLimiter(inflowRate, bucketSize)

	return func(c *gin.Context) {
		if rateLimiter.IsAllowed() {
			c.Next()
		} else {
			message := "Request is rate limited, please try again later"
			log.Println(message)

			response := models.Response{Message: message}

			c.AbortWithStatusJSON(http.StatusTooManyRequests, response)
			return
		}
	}
}
