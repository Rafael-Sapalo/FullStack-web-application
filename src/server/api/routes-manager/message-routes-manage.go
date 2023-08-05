package routes_manager

import "github.com/gin-gonic/gin"

func SetMessageRoutes(msgRoute *gin.RouterGroup) {
	msgRoute.GET("")
	msgRoute.GET("/:conversationID")
}
