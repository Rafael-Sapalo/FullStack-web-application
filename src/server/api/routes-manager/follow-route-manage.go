package routes_manager

import "github.com/gin-gonic/gin"

func SetFollowRoutes(followRoutes *gin.RouterGroup) {
	followRoutes.POST("" /*hdl follow user*/)
	followRoutes.DELETE("" /*hdl unfollow*/)
	followRoutes.GET("/followers" /*hdl get followers*/)
	followRoutes.GET("/following" /*hdl get all following*/)
}
