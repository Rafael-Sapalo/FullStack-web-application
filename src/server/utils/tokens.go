package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte("bjqooe4nky2i28e1ugehwbom11oyv6erce8");

func IsAuthenticated(c *gin.Context)  {
	var tokenStr = c.GetHeader("Authorization");

	if tokenStr == "" {
		c.JSON(401, gin.H{"error": "Unauthorized"});
		c.Abort();
		return;
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error)  {
		return secretKey, nil;
	})
	if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "Unauthorized"});
		c.Abort();
		return;
	}

	c.Next();
}