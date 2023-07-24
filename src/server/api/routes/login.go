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
	Logged ,LogErr := LoginUser(ctx, Logs.Email, Logs.Password);
	var UserId, UserErr = GetUserID(ctx, Logs.Email);
	if UserErr != nil {
		ctx.JSON(UserErr.Code, gin.H{"msg": UserErr.Message})
		return;
	}
	if LogErr != nil && Logged == false{
		ctx.JSON(LogErr.Code, gin.H{"msg": LogErr.Message});
	}
	if Logged {
		session := sessions.Default(ctx);
		session.Set("user_id", UserId)
		session.Save()
		
		var token, TokErr = utils.GenerateToken(UserId);
		if TokErr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
			return;
		}
		ctx.JSON(utils.SuccessfullyLoggedIn.Code, gin.H{"token": token})
	}
}