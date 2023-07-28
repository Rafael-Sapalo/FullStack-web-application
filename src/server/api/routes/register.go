package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/Rafael-Sapalo/FullStack-web-application/server/utils"
)

func GetUserID(ctx *gin.Context, email string) (int, *utils.ErrorMessage) {
	var db = ctx.MustGet("db").(*sql.DB)
	var trans, TrErr = db.Begin()
	var cmd = "SELECT id FROM users WHERE email = ?"
	var UserId int
	if TrErr != nil {
		return 0, utils.ErrInternalServerError
	}
	defer trans.Rollback()
	if err := trans.QueryRow(cmd, email).Scan(&UserId); err != nil {
		return 0, utils.ErrInternalServerError
	}
	if err := trans.Commit(); err != nil {
		return 0, utils.ErrorCommit
	}
	return UserId, nil
}

func RegisterRoute(ctx *gin.Context) {
	var userData utils.UserData

	if err := ctx.ShouldBindJSON(&userData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userData.Email == "" || userData.Password == "" || userData.Username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing parameters"})
		return
	}
	if Regerr := RegisterUser(ctx, userData); Regerr != nil {
		ctx.JSON(Regerr.Code, gin.H{"msg": Regerr.Message})
		return
	}

	var UserId, Regerr = GetUserID(ctx, userData.Email)
	if Regerr != nil {
		ctx.JSON(Regerr.Code, gin.H{"msg": Regerr.Message})
		return
	}
	session := sessions.Default(ctx)
	session.Set("user_id", UserId)
	session.Save()
	fmt.Println(UserId)
	ctx.Set("user", userData)
	ctx.JSON(utils.SuccessfullyRegistered.Code, gin.H{"Success": utils.SuccessfullyRegistered.Message})
}
