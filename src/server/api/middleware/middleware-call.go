package middleware

import "github.com/gin-gonic/gin"

func RegisterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next();
	}
}
