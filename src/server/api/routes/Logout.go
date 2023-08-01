package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LogoutRoute(ctx *gin.Context) {
	var session = sessions.Default(ctx)
	session.Clear()
	session.Save()
	ctx.JSON(http.StatusOK, gin.H{"Success": "Successfully  logout"})
}
