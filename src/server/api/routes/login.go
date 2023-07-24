package routes

import (
	"net/http"
	"time"

	"github.com/Rafael-Sapalo/FullStack-web-application/server/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte("bjqooe4nky2i28e1ugehwbom11oyv6erce8");

func LoginRoute(c *gin.Context) {
	userID := 123;
	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenStr, err := token.SignedString(secretKey);
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
	}
	session := sessions.Default(c);
	session.Set("username", "Rafael");
	session.Save();

	c.Header("Authorization", tokenStr);
	c.JSON(http.StatusOK, gin.H{"status": "login successful", "token": tokenStr});
}

func LogoutRoute(ctx *gin.Context) {
	var Logs utils.UserData;
	if err := ctx.ShouldBindJSON(&Logs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if Logs.Email == "" || Logs.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing params"})
		return
	}
	session := sessions.Default(ctx);
	session.Clear();
	session.Save();

	ctx.JSON(http.StatusOK, gin.H{"status": "logout successful"});
}