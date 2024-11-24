package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func HomePage(ctx *gin.Context) {

	isAuthenticated := ctx.GetBool("isAuthenticated")
	if isAuthenticated {
		ctx.HTML(http.StatusOK, "home.html", gin.H{
			"links": []string{"appointment", "profile", "logout"},
		})
	} else {
		ctx.HTML(http.StatusOK, "home.html", gin.H{
			"links": []string{"signup", "login"},
		})
	}
}
