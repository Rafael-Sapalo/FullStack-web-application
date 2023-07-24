package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte("bjqooe4nky2i28e1ugehwbom11oyv6erce8");

func GenerateToken(userID int) (string, error) {
	var Exp = time.Now().Add(24 * time.Hour);
	var claim = jwt.MapClaims{
		"user_id": userID,
		"exp": Exp.Unix(),
	}
	var tokenHelper = jwt.NewWithClaims(jwt.SigningMethodES256, claim);
	var token, err = tokenHelper.SignedString([]byte(secretKey));
	if err != nil {
		return "", fmt.Errorf("Cannot generate token %w", err);
	}
	return token, nil;
}

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