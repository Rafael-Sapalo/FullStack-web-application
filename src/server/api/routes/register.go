package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Rafael-Sapalo/FullStack-web-application/server/utils"
)

func RegisterRoute(c *gin.Context) {

	var userData utils.UserData;
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()});
		return;
	}
	if userData.Email == "" || userData.Password == "" || userData.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing parameters"});
		return;
	}
	userData.Password,_ = utils.HashPass(userData.Password)
	fmt.Println(userData);
	if err := utils.RegisterUser(userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()});
		return;
	}
	c.JSON(http.StatusCreated, gin.H{"status": "register successful"});
}