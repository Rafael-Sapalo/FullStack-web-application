package main

import (
	"database/sql"
	"net/http"

	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/middleware"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/routes"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/config"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte("your-secret-key")

func attachDB(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func main() {

	router := gin.Default()
	db, err := config.ConnectDB()

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router.Use(attachDB(db))
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Path: "/src/server",
		MaxAge: 60 * 60,
		HttpOnly: true,
		Secure: true,
		SameSite: http.SameSiteDefaultMode,
	})
	router.Use(sessions.Sessions("UserSession", store))

	router.GET("/", middleware.RateLimitIndex(), routes.IndexRoute)
	router.GET("/get-all-users", routes.GetAllUsersRoute)

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", middleware.RegisterMiddleware(), routes.RegisterRoute)
		authRoutes.POST("/login", routes.LoginRoute)
		authRoutes.POST("/logout", routes.LogoutRoute);
		authRoutes.GET("/protected", func(ctx *gin.Context) {
			session := sessions.Default(ctx);
			userId := session.Get("user_id");
			if userId == nil {
				ctx.JSON(utils.ErrorUnauthorized.Code, gin.H{"error": "Unauthorized"});
				return;
			}
			ctx.JSON(http.StatusOK, gin.H{"message": userId});
		})
	}
	userRoutes := router.Group("/users/:username")
	userRoutes.Use(middleware.Authenticate())
	{
		userRoutes.GET("" /*hdl profile*/);//Todo: add basic user route
		userRoutes.PUT("/edit-profile");//Todo: add edit-profile route
	}
	postRoutes := router.Group("/posts")
	postRoutes.Use(middleware.Authenticate())
	{
		postRoutes.GET("" /*hdl get all posts*/)
		postRoutes.GET("/:postId" /*hdl post by id*/)
		postRoutes.POST("" /*hdl create post*/)
		postRoutes.PUT("/:postId" /*hdl edit post*/)
		postRoutes.DELETE("/:postId" /*hdl delete post*/)
	}
	commentRoutes := router.Group("/posts/:postId/comments")
	commentRoutes.Use(middleware.Authenticate())
	{
		commentRoutes.GET("" /*hdl get all comments*/)
		commentRoutes.POST("" /*hdl create comment*/)
		commentRoutes.PUT("/:commentId" /*hdl edit comment*/)
		commentRoutes.DELETE("/:commentId" /*hdl delete comment*/)
	}
	likeRoutes := router.Group("/like")
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

	router.Run(":8080")
}
