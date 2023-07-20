package main

import (
	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/routes"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("../public/*.html")

	router.GET("/", 
		middleware.RegiserMiddleware(), 
		middleware.RateLimitIndex(), 
		routes.IndexRoute,
	)
	router.POST("/register", middleware.RateLimitRegister(), routes.RegisterRoute)
	router.GET("/login", routes.LoginRoute)

	router.Run(":8080")
}
