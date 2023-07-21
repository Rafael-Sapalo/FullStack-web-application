package main

import (
	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/middleware"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/routes"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/config"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte("your-secret-key");


func main() {

	router := gin.Default()
	store := cookie.NewStore([]byte("secret"));
	router.Use(sessions.Sessions("usersession", store));

	router.GET("/", middleware.RateLimitIndex(), routes.IndexRoute)

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", middleware.RegisterMiddleware(), routes.RegisterRoute)
		authRoutes.POST("/login", routes.LoginRoute)
		authRoutes.POST("/logout", routes.LogoutRoute) //TODO: ADD logout route
		authRoutes.GET("/protected", utils.IsAuthenticated, func(ctx *gin.Context) {
			userID := ctx.MustGet("userID").(int);
			ctx.JSON(200, gin.H{"message": "This is a protected route", "userID": userID});
		})
	}
	userRoutes := router.Group("/users/:username")
	{
		userRoutes.GET("" /*hdl profile*/)
		userRoutes.PUT("/edit-profile")
	}
	postRoutes := router.Group("/posts")
	{
		postRoutes.GET("" /*hdl get all posts*/)
		postRoutes.GET("/:postId" /*hdl post by id*/)
		postRoutes.POST("" /*hdl create post*/)
		postRoutes.PUT("/:postId" /*hdl edit post*/)
		postRoutes.DELETE("/:postId" /*hdl delete post*/)
	}
	commentRoutes := router.Group("/posts/:postId/comments")
	{
		commentRoutes.GET("" /*hdl get all comments*/)
		commentRoutes.POST("" /*hdl create comment*/)
		commentRoutes.PUT("/:commentId" /*hdl edit comment*/)
		commentRoutes.DELETE("/:commentId" /*hdl delete comment*/)
	}
	likeRoutes := router.Group("/like")
	{
		likeRoutes.POST("/:postId" /*hdl create like*/)
		likeRoutes.DELETE("/:postId" /*hdl remove like*/)
	}
	var followRoutes = router.Group("/follow/:userId")
	{
		followRoutes.POST("" /*hdl follow user*/)
		followRoutes.DELETE("" /*hdl unfollow*/)
		followRoutes.GET("/followers" /*hdl get followers*/)
		followRoutes.GET("/following" /*hdl get all following*/)
	}
	var searchRoute = router.Group("/search")
	{
		searchRoute.GET("" /*hdl search*/)
	}
	var exploreRoutes = router.Group("/explore")
	{
		exploreRoutes.GET("" /*hdl explore*/)
	}
	var messageRoutes = router.Group("/messages")
	{
		messageRoutes.GET("" /*hdl get all message*/)
		messageRoutes.GET("/:conversationId" /*get conv by id*/)
	}

	db, err := config.ConnectDB()

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router.Run(":8080")
}
