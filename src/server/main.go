package main

import (
	"net/http"

	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/middleware"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/routes"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/config"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	var router = gin.Default()
	var db, err = config.ConnectDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var store = cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteDefaultMode,
	})

	router.Use(config.AttachDB(db))
	router.Use(sessions.Sessions("basic", store))

	router.GET("/", middleware.RateLimitIndex(), routes.IndexRoute)
	router.GET("/get-all-users", routes.GetAllUsersRoute)

	var authRoutes = router.Group("/auth")
	authRoutes.Use(middleware.RateLimit())
	{
		authRoutes.POST("/register", middleware.RegisterMiddleware(), routes.RegisterRoute)
		authRoutes.POST("/login", routes.LoginRoute)
		authRoutes.GET("/logout", routes.LogoutRoute)
	}
	var protectedRoutes = router.Group("/protected")
	protectedRoutes.Use(middleware.IsAdmin, middleware.Authenticate(), middleware.RateLimit())
	{
		protectedRoutes.GET("/check-protection", func(ctx *gin.Context) {
			var session = sessions.Default(ctx)
			var userID = session.Get("user_id")
			if userID == nil {
				ctx.JSON(utils.ErrorUnauthorized.Code, gin.H{"error": utils.ErrorUnauthorized.Message})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": userID})
		})
		protectedRoutes.GET("/get-all-users", routes.GetAllUsersRoute)
	}
	var userRoutes = router.Group("/users/:username")
	userRoutes.Use(middleware.Authenticate())
	{
		userRoutes.GET("" /*hdl profile*/) //Todo: add basic user route
		userRoutes.PUT("/edit-profile")    //Todo: add edit-profile route
	}
	var postRoutes = router.Group("/posts")
	postRoutes.Use(middleware.Authenticate())
	{
		postRoutes.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "you have auth"})
		} /*hdl get all posts*/)
		postRoutes.GET("/:postId" /*hdl post by id*/)
		postRoutes.POST("" /*hdl create post*/)
		postRoutes.PUT("/:postId" /*hdl edit post*/)
		postRoutes.DELETE("/:postId" /*hdl delete post*/)
	}
	var commentRoutes = router.Group("/posts/:postId/comments")
	commentRoutes.Use(middleware.Authenticate())
	{
		commentRoutes.GET("" /*hdl get all comments*/)
		commentRoutes.POST("" /*hdl create comment*/)
		commentRoutes.PUT("/:commentId" /*hdl edit comment*/)
		commentRoutes.DELETE("/:commentId" /*hdl delete comment*/)
	}
	var likeRoutes = router.Group("/like")
	likeRoutes.Use(middleware.Authenticate())
	{
		likeRoutes.POST("/:postId" /*hdl create like*/)
		likeRoutes.DELETE("/:postId" /*hdl remove like*/)
	}
	var followRoutes = router.Group("/follow/:userId")
	followRoutes.Use(middleware.Authenticate())
	{
		followRoutes.POST("" /*hdl follow user*/)
		followRoutes.DELETE("" /*hdl unfollow*/)
		followRoutes.GET("/followers" /*hdl get followers*/)
		followRoutes.GET("/following" /*hdl get all following*/)
	}
	var searchRoute = router.Group("/search")
	searchRoute.Use(middleware.Authenticate())
	{
		searchRoute.GET("" /*hdl search*/)
	}
	var exploreRoutes = router.Group("/explore")
	exploreRoutes.Use(middleware.Authenticate())
	{
		exploreRoutes.GET("" /*hdl explore*/)
	}
	var messageRoutes = router.Group("/messages")
	messageRoutes.Use(middleware.Authenticate())
	{
		messageRoutes.GET("" /*hdl get all message*/)
		messageRoutes.GET("/:conversationId" /*get conv by id*/)
	}

	var RunErr = router.Run(":8080")
	if RunErr != nil {
		return
	}
}
