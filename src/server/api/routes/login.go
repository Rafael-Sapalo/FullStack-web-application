package routes

import (
	"fmt"
	"net/http"

	"github.com/Rafael-Sapalo/FullStack-web-application/server/utils"
	"github.com/gin-gonic/gin"
)

func LoginRoute(c *gin.Context) {
	var userData utils.UserData;
	fmt.Println("Trying to login...")
	c.JSON(http.StatusCreated, gin.H{"status": "trying to login..."})
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()});
		return;
	}
	c.JSON(http.StatusCreated, gin.H{"status": "login successful"});
}
