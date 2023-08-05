package routes_manager

import (
	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/middleware"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetSearchRoutes(searchGrp *gin.RouterGroup) {
	searchGrp.GET("")
}

func SetExploreRoute(explRoutes *gin.RouterGroup) {
	explRoutes.GET("")
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
		SetFollowRoutes(userRoutes.Group("/follow/:username"))
		SetSearchRoutes(userRoutes.Group("/search"))
		SetExploreRoute(userRoutes.Group("/explore"))
		SetMessageRoutes(userRoutes.Group("/messages"))
	}
}
