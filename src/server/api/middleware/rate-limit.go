package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"net/http"
)

var limiter = rate.NewLimiter(rate.Limit(5), 15)

var registerLimiter = rate.NewLimiter(rate.Limit(5), 5)

func RateLimitIndex() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !limiter.Allow() {
			ctx.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Too many unsuccessful requests ",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func RateLimitRegister() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := registerLimiter.Wait(ctx); err != nil {
			ctx.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Too many unsuccessful requests",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func RateLimit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := limiter.Wait(ctx); err != nil {
			ctx.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Too many request sent",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
