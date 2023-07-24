package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var session = sessions.Default(ctx);
		var userID = session.Get("user_id");
		if userID == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unhautorized"})
			ctx.Abort();
			return;
		}
		ctx.Next();
	}
}