package middleware

import (
	"github.com/Rafael-Sapalo/FullStack-web-application/server/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"net/http"
)

var limiter *rate.Limiter = rate.NewLimiter(rate.Limit(20), 150)
var LikesLimit *rate.Limiter = rate.NewLimiter(rate.Limit(utils.NbLikeRequest), utils.BurstLikeRequest)
var registerLimiter *rate.Limiter = rate.NewLimiter(rate.Limit(20), 150)

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

func LikesRateLimit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := LikesLimit.Wait(ctx); err != nil {
			ctx.JSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
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
