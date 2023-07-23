package routes

import (
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
	if Regerr := RegisterUser(c, userData.Email, userData.Password, userData.Username); Regerr != nil {
		c.JSON(Regerr.Code, gin.H{"msg": Regerr.Message});
		return;
	}
}