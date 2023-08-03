package routes_manager

import (
	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/middleware"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetPostsRoutes(postRoutes *gin.RouterGroup) {
	postRoutes.POST("")
	postRoutes.GET("/:postID")
	postRoutes.PUT("/:postID")
	postRoutes.DELETE("/:postID")
}

func SetUserRoutes(router *gin.Engine) {
	var userRoutes = router.Group("/user/:username")
	userRoutes.Use(middleware.Authenticate(), middleware.RateLimit())
	{
		userRoutes.GET("/try", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": utils.SuccessfullyLoggedIn.Message})
		})
		userRoutes.GET("")
		userRoutes.PUT("/edit-profile")
		SetPostsRoutes(userRoutes.Group("/posts"))

	}
}
