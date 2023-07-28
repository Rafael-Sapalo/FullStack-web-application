package routes

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int
	Username string
	Email    string
	password string
}

func GetAllUsersRoute(ctx *gin.Context) {
	db := ctx.MustGet("db").(*sql.DB)
	var statement = "SELECT id, username, email, password FROM users"
	rows, err := db.Query(statement)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}
	defer rows.Close()
	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
			return
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "users": users})

}
