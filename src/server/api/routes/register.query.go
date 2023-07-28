package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/Rafael-Sapalo/FullStack-web-application/server/utils"
)

func RegisterUser(ctx *gin.Context, userData utils.UserData) *utils.ErrorMessage {
	var db = ctx.MustGet("db").(*sql.DB)
	var MailExist, MailErr = utils.IsEmailAlreadyExist(db, userData.Email)
	var UserExist, UserErr = utils.IsUsernameAlreadyExist(db, userData.Username)
	var hash, err = utils.HashPass(userData.Password)

	if err != nil {
		return utils.ErrorHashingPassword
	}
	if MailErr != nil || UserErr != nil {
		return utils.ErrInternalServerError
	}
	if MailExist || UserExist {
		return utils.ErrorEmailAlrdExists
	}
	var trans, Terr = db.Begin()
	if Terr != nil {
		return utils.ErrInternalServerError
	}
	if userData.Role != "admin" {
		var cmd = "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
		if _, err := trans.Exec(cmd, userData.Username, userData.Email, hash); err != nil {
			trans.Rollback()
			return utils.ErrorInsertingUserData
		}
		if err := trans.Commit(); err != nil {
			return utils.ErrorCommit
		}
	} else {
		var cmd = "INSERT INTO users (username, email, password, roles) VALUES (?, ?, ?, ?)"
		if _, err := trans.Exec(cmd, userData.Username, userData.Email, hash, userData.Role); err != nil {
			trans.Rollback()
			return utils.ErrorInsertingUserData
		}
		if err := trans.Commit(); err != nil {
			return utils.ErrorCommit
		}
	}
	return nil
}
