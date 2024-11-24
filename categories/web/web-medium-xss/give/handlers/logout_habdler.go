package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func LogoutPage(ctx *gin.Context) {
	ctx.SetCookie("Authorization", "", -1, "/", "", false, true)
	ctx.Set("isAuthenticated", false)
	ctx.Redirect(http.StatusSeeOther, "/home/")
}
