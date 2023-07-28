package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var session = sessions.Default(ctx)
		var userID = session.Get("user_id")
		if userID == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func IsAdmin(ctx *gin.Context) {
	var session = sessions.Default(ctx)
	var id = session.Get("user_id")
	fmt.Println(id)
	if id == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}

	ctx.Next()
}
