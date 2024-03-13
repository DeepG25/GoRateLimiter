package models

import (
	"log"
	"time"
)

type RateLimiter interface {
	IsAllowed() bool
}

type TokenBucketRateLimiter struct {
	InflowRate             int
	BucketSize             int
	TokenCount             int
	LastRefillTime         time.Time
	RefreshWindowInSeconds int
}

func (rateLimiter *TokenBucketRateLimiter) Refill() {

	currentTime := time.Now().UTC()
	elapsedTime := currentTime.Sub(rateLimiter.LastRefillTime).Seconds() / float64(rateLimiter.RefreshWindowInSeconds)
	if elapsedTime < 1 {
		return
	}

	rateLimiter.TokenCount = min(rateLimiter.BucketSize, rateLimiter.TokenCount+rateLimiter.InflowRate*int(elapsedTime))
	rateLimiter.LastRefillTime = currentTime
	log.Println("RateLimiter Log: Bucket refilled")
}

func (rateLimiter *TokenBucketRateLimiter) IsAllowed() bool {
	rateLimiter.Refill()
	log.Printf("%v tokens left in bucket\n", rateLimiter.TokenCount)

	if rateLimiter.TokenCount != 0 {
		rateLimiter.TokenCount = rateLimiter.TokenCount - 1
		return true
	}

	return false
}
