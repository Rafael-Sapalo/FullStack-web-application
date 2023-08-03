package routes_manager

import (
	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/middleware"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/routes"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetProtectedRoutes(router *gin.Engine) {
	var protectedRoutes = router.Group("/protected")
	protectedRoutes.Use(middleware.Authenticate())
	{
		protectedRoutes.GET("/check-protection", func(ctx *gin.Context) {
			var session = sessions.Default(ctx)
			var userID = session.Get("user_id")
			if userID == nil {
				ctx.JSON(utils.ErrorUnauthorized.Code, gin.H{"error": utils.ErrorUnauthorized.Message})
			}
			ctx.JSON(http.StatusOK, gin.H{"msg": userID})
		})
		protectedRoutes.GET("/get-all-users", routes.GetAllUsersRoute)
	}
}
