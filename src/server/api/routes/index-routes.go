package routes

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
)

func IndexRoute(c *gin.Context) {
	fmt.Println("Connecting to server...");
	c.JSON(http.StatusOK, gin.H{"status": "connected to server"});
}