package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexRoute(ctx *gin.Context) {
	fmt.Println("Connecting to server...")
	tmpl, err := template.ParseFiles("../../../public/index.html")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if err := tmpl.Execute(ctx.Writer, nil); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
}
