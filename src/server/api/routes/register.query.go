package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/Rafael-Sapalo/FullStack-web-application/server/utils"
)

func RegisterUser(c *gin.Context, email string, password string, username string) *utils.ErrorMessage {
	var hash, err = utils.HashPass(password);
	if err != nil {
		return utils.ErrorHashingPassword;
	}
	var db = c.MustGet("db").(*sql.DB);
	var exist, Merr = utils.IsEmailAlreadyExist(db, email);
	if Merr != nil {
		return utils.ErrInternalServerError;
	}
	if exist {
		return utils.ErrorEmailAlrdExists;
	}
	var trans, Terr = db.Begin();
	if Terr != nil {
		return utils.ErrInternalServerError
	}
	var cmd = "INSERT INTO users (username, email, password) VALUES (?, ?, ?)";
	if _, err := trans.Exec(cmd, username, email, hash); err != nil {
		trans.Rollback()
		return utils.ErrorInsertingUserData;
	}
	if err := trans.Commit(); err != nil {
		return utils.ErrorCommit
	}
	return utils.SuccessfullyRegistered;
}
