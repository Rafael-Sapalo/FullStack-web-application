package routes_manager

import (
	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/middleware"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/routes"
	"github.com/gin-gonic/gin"
)

func SetAuthRoutes(router *gin.Engine) {
	var authRoutes = router.Group("/auth")
	authRoutes.Use(middleware.RateLimit())
	{
		authRoutes.POST("/register", middleware.RegisterMiddleware(), routes.RegisterRoute)
		authRoutes.POST("/login", routes.LoginRoute)
		authRoutes.GET("/logout", routes.LogoutRoute)
	}
}
