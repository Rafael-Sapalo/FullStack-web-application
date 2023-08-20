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

func HandleUsrIdInCtx(ctx *gin.Context) bool {
	userIDval, exist := ctx.Get("user_id")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found in the context"})
		return false
	}
	userID, ok := userIDval.(int)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user_id has unexpected type"})
		return false
	}
	fmt.Println(userID)
	return true
}

func IsAdmin(ctx *gin.Context) {
	var session = sessions.Default(ctx)
	var id = session.Get("user_id")
	fmt.Println(id)
	if id == nil {
		return
	} else {
		//var id, _ = ctx.Get("user_id")
		if id == nil {
			fmt.Println("id is nil")
		}
		fmt.Println(id)
		if id == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
