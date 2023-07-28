package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/Rafael-Sapalo/FullStack-web-application/server/utils"
)

func GetPassword(db *sql.DB, email string) (string, *utils.ErrorMessage) {
	var trans, TrErr = db.Begin()
	var cmd = "SELECT password FROM users WHERE email = ?"
	var password string
	if TrErr != nil {
		return "", utils.ErrInternalServerError
	}
	defer trans.Rollback()
	if err := trans.QueryRow(cmd, email).Scan(&password); err != nil {
		return "", utils.ErrInternalServerError
	}
	if err := trans.Commit(); err != nil {
		return "", utils.ErrorCommit
	}
	return password, nil
}

func LoginUser(ctx *gin.Context, email string, password string) (bool, *utils.ErrorMessage) {
	var db = ctx.MustGet("db").(*sql.DB)
	var MailExist, MailErr = utils.IsEmailAlreadyExist(db, email)
	var dbPassword, PassErr = GetPassword(db, email)

	if MailErr != nil {
		return false, utils.ErrInternalServerError
	}
	if MailExist {
		if PassErr != nil {
			return false, PassErr
		}
		var isGoodPassword = utils.CmpHash(dbPassword, password)
		if isGoodPassword {
			return true, utils.IsGoodPassword
		}
		return false, utils.ErrInternalServerError
	}
	return false, utils.ErrInternalServerError
}
