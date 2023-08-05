package routes_manager

import (
	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/middleware"
	"github.com/gin-gonic/gin"
)

func SetPostsRoutes(postRoutes *gin.RouterGroup) {
	postRoutes.POST("")
	postRoutes.GET("/:postID")
	postRoutes.PUT("/:postID")
	postRoutes.DELETE("/:postID")
	SetCommentRoutes(postRoutes.Group("/:postID/comments"))
}

func SetCommentRoutes(comRoutes *gin.RouterGroup) {
	comRoutes.GET("")
	comRoutes.POST("")
	comRoutes.PUT("/:commentID")
	comRoutes.DELETE("/:commentID")
}

func SetLikesRoutes(likeRoutes *gin.RouterGroup) {
	likeRoutes.Use(middleware.LikesRateLimit())
	likeRoutes.POST("/:postID")
	likeRoutes.DELETE("/:posID")
}
