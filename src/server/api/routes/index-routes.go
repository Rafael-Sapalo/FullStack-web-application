package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexRoute(ctx *gin.Context) {
	fmt.Println("Connecting to server...")
	ctx.JSON(http.StatusOK, gin.H{"msg": "Welcome to the webapp"})
}
