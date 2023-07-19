package routes

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func LoginRoute(c *gin.Context) {
	fmt.Println("Trying to login...")
	c.JSON(200, gin.H{
		"message": "call form login page",
	})
}
