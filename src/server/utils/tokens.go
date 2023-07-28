package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte("dkmohiqfoplzzznjdqorjnaoupgicjrurhs")

func GenerateToken(userID int) (string, error) {
	var Exp = time.Now().Add(24 * time.Hour)
	var claim = jwt.MapClaims{
		"user_id": userID,
		"exp":     Exp.Unix(),
	}
	var tokenHelper = jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	var token, err = tokenHelper.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("Cannot generate token %w", err)
	}
	return token, nil
}

func IsAuthenticated(ctx *gin.Context) {
	var tokenStr = ctx.GetHeader("Authorization")

	if tokenStr == "" {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}

	ctx.Next()
}
