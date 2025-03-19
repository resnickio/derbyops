// Package middleware provides HTTP middleware functions for the API.
package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/resnickio/derbyops/api/logger"
	"golang.org/x/time/rate"
)

var (
	// ipLimiter stores rate limiters for each IP address
	ipLimiter = make(map[string]*rate.Limiter)
	// lastAccess tracks when each limiter was last used
	lastAccess = make(map[string]time.Time)
	mu        sync.RWMutex
)

// RateLimit creates a middleware for rate limiting requests based on IP address.
// It uses a token bucket algorithm with the specified requests per minute.
func RateLimit(requestsPerMinute int) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := getLimiter(ip, requestsPerMinute)

		if !limiter.Allow() {
			logger.Warn("Rate limit exceeded", map[string]interface{}{
				"ip":      ip,
				"path":    c.Request.URL.Path,
				"method":  c.Request.Method,
				"rpmMax": requestsPerMinute,
			})

			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests. Please try again later.",
			})
			c.Abort()
			return
		}

		// Update last access time
		mu.Lock()
		lastAccess[ip] = time.Now()
		mu.Unlock()

		c.Next()
	}
}

// getLimiter returns a rate limiter for the specified IP address.
// If a limiter doesn't exist for the IP, it creates a new one.
func getLimiter(ip string, requestsPerMinute int) *rate.Limiter {
	mu.RLock()
	limiter, exists := ipLimiter[ip]
	mu.RUnlock()

	if !exists {
		mu.Lock()
		// Double check after acquiring write lock
		limiter, exists = ipLimiter[ip]
		if !exists {
			limiter = rate.NewLimiter(rate.Every(time.Minute/time.Duration(requestsPerMinute)), requestsPerMinute)
			ipLimiter[ip] = limiter
			lastAccess[ip] = time.Now()
		}
		mu.Unlock()
	}

	return limiter
}

// cleanupLimiters periodically removes old limiters to prevent memory leaks.
// This should be called in a goroutine when the server starts.
func cleanupLimiters() {
	ticker := time.NewTicker(10 * time.Minute)
	for range ticker.C {
		mu.Lock()
		now := time.Now()
		for ip := range ipLimiter {
			if lastAccess[ip].Add(time.Hour).Before(now) {
				delete(ipLimiter, ip)
				delete(lastAccess, ip)
			}
		}
		mu.Unlock()
	}
} 