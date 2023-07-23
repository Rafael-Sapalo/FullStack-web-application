package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/Rafael-Sapalo/FullStack-web-application/server/utils"
)

func LoginUser(c *gin.Context, email string, password string) *utils.ErrorMessage {
	var db = c.MustGet("db").(*sql.DB);
}