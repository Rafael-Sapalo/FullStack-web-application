package main

import (
	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/routes"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/middleware"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/config"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	db, err := config.ConnectDB();

	if err != nil {
		panic(err.Error());
	}
	defer db.Close();
	router.GET("/",middleware.RateLimitIndex(), routes.IndexRoute);

	//!Auth routes
	router.POST("/register", middleware.RegisterMiddleware(), routes.RegisterRoute);
	router.POST("/login", routes.LoginRoute);
	router.POST("/logout"); //!TODO: Add logout route

	//!User routes
	router.GET("/profile/:username", /*Add profile route*/);
	router.PUT("/edit-profile/:username", /*Add edit profile route*/);
	
	//!Post routes
	router.GET("/posts", /*Add get all posts route*/);
	router.GET("/posts/:postId", /*Add get post by username route*/);
	router.POST("create-post", /*Add create post route*/);
	router.PUT("/edit-post/:postId", /*Add edit post route*/);
	router.DELETE("/delete-post/:postId", /*Add delete post route*/);

	//!Comment routes
	router.GET("/comments/:postId", /*Add get all comments route*/);
	router.POST("/create-comment/:postId", /*Add create comment route*/);
	router.PUT("/edit-comment/:commentId", /*Add edit comment route*/);
	router.DELETE("/delete-comment/:commentId", /*Add delete comment route*/);

	//!Like routes
	router.POST("/create-like/:postId", /*Add create like route*/);
	router.DELETE("/delete-like/:postId", /*Add delete like route*/);

	//!Follow routes
	router.POST("/follow/:userId", /*Add follow route*/);
	router.DELETE("/unfollow/:userId", /*Add unfollow route*/);
	router.GET("/followers/:userId", /*Add get followers route*/);
	router.GET("/following/:userId", /*Add get following route*/);

	//!Search routes
	router.GET("/search", /*Add search route*/);

	//!explore
	router.GET("/explore", /*Add explore route*/);

	//!message routes
	router.GET("/messages", /*Add get all messages route*/);
	router.GET("/messages/:conversationId", /*Add get message by id route*/);

	router.Run(":8080")
}
