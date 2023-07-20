package main

import (
	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/routes"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/middleware"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/config"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	db, err := config.ConnectDB();

	if err != nil {
		panic(err.Error());
	}
	defer db.Close();

	router.GET("/", 
		middleware.RegisterMiddleware(), 
		middleware.RateLimitIndex(), 
		routes.IndexRoute,
	)
	router.POST("/register", routes.RegisterRoute)
	router.GET("/login", routes.LoginRoute)

	router.Run(":8080")
}
