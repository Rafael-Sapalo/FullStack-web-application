package routes_manager

import (
	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/middleware"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetAuthRoutes(router *gin.Engine) {
	var authRoutes = router.Group("/auth")
	authRoutes.Use(middleware.RateLimit())
	{
		authRoutes.GET("/try", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "I'm just here to try some features"})
		})
		authRoutes.POST("/register", middleware.RegisterMiddleware(), routes.RegisterRoute)
		authRoutes.POST("/login", routes.LoginRoute)
		authRoutes.GET("/logout", routes.LogoutRoute)
	}
}
