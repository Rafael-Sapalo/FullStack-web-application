package main

import (
	routes_manager "github.com/Rafael-Sapalo/FullStack-web-application/server/api/routes-manager"
	"net/http"

	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/middleware"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/routes"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/config"
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

	routes_manager.SetAuthRoutes(router)
	routes_manager.SetProtectedRoutes(router)
	routes_manager.SetUserRoutes(router)

	router.GET("/", middleware.RateLimitIndex(), routes.IndexRoute)
	router.GET("/get-all-users", routes.GetAllUsersRoute)

	var RunErr = router.Run(":8080")
	if RunErr != nil {
		return
	}
}
