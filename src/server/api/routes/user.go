package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int;
	Username string;
}

func GetAllUsersRoute(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB);
	var statement = "SELECT id, username FROM users" ;
	rows, err := db.Query(statement);
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"});
		return;
	}
	defer rows.Close();
	var users []User;

	for rows.Next() {
		var user User;
		err := rows.Scan(&user.Id, &user.Username);
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error"});
			return;
		}
		users = append(users, user);
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"});
		return;
	}
	fmt.Printf("this is the is %d\n", users[0].Id);
	c.JSON(http.StatusOK, gin.H{"status": "success", "users": users});

}