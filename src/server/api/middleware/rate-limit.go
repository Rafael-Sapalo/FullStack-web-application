package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"net/http"
)

var limiter = rate.NewLimiter(rate.Limit(5), 5);

func RateLimitIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Too many unsuc requests ",
			})
			c.Abort()
			return
		}
		c.Next();
	}
}

func RateLimitRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := limiter.Wait(c); err != nil {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message" : "Too many unsuccessful requests",
			})
			c.Abort();
			return;
		}
		c.Next();
	}
}
