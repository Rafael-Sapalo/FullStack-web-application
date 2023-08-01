package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Rafael-Sapalo/FullStack-web-application/server/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte("bjqooe4nky2i28e1ugehwbom11oyv6erce8")

func GetUsersRole(ctx *gin.Context, email string, id int) (string, *utils.ErrorMessage) {
	var db = ctx.MustGet("db").(*sql.DB)
	var trans, TrErr = db.Begin()
	var cmd = "SELECT roles FROM users WHERE id = ?"
	var role string
	if TrErr != nil {
		return "", utils.ErrInternalServerError
	}
	defer trans.Rollback()
	if err := trans.QueryRow(cmd, id).Scan(&role); err != nil {
		return "", utils.ErrInternalServerError
	}
	if err := trans.Commit(); err != nil {
		return "", utils.ErrorCommit
	}
	return role, nil
}

func LoginRoute(ctx *gin.Context) {
	var Logs utils.UserData
	if err := ctx.ShouldBindJSON(&Logs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if Logs.Email == "" || Logs.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing params"})
		return
	}
	var Logged, LogErr = LoginUser(ctx, Logs.Email, Logs.Password)
	var UserId, UserErr = GetUserID(ctx, Logs.Email)
	if UserErr != nil {
		ctx.JSON(UserErr.Code, gin.H{"msg": UserErr.Message})
		return
	}
	if LogErr != nil && !Logged {
		ctx.JSON(LogErr.Code, gin.H{"msg": LogErr.Message})
	}
	if Logged {
		var session = sessions.Default(ctx)
		session.Set("user_id", UserId)
		var err = session.Save()
		if err != nil {
			return
		}
		var token, TokErr = utils.GenerateToken(UserId)
		if TokErr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
			return
		}
		fmt.Printf("This is the id logged in %d\n", UserId)
		ctx.JSON(utils.SuccessfullyLoggedIn.Code, gin.H{"token": token})
	}
}
