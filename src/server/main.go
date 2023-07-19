package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/routes"
)

func GetData(c *gin.Context) {
	name := c.Param("name");
	c.JSON(200, gin.H{
		"message": "Hello World! " + name,
	})
}

func main() {

	router := gin.Default();

	router.GET("/:name", GetData);
	router.GET("/login", routes.LoginRoute);

	router.Run(":8080");
}
