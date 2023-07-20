package routes

import (
	"net/http"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/utils"
	"github.com/gin-gonic/gin"
	"fmt"
)

func RegisterRoute(c *gin.Context) {
	var userData utils.UserData;
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()});
		return;
	}
	fmt.Printf("name = %s, fname = %s, email = %s, passwoord = %s\n", userData.Name, userData.FirstName, userData.Email, userData.Password);
	c.JSON(http.StatusCreated, gin.H{"status": "register successful"});
}