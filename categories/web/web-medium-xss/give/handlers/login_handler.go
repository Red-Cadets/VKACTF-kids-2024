package handlers

import (
	"example/v3/auth"
	"example/v3/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func LoginPage(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "login.html", gin.H{
		"Login": "Login Page",
	})
}
func LoginHandler(ctx *gin.Context) {

	login := ctx.PostForm("Login")
	password := ctx.PostForm("Password")

	user, err := auth.AuthenticateUser(login, password)
	if err != nil {
		ctx.HTML(http.StatusOK, "login.html", gin.H{
			"Login":        "Login Page",
			"ErrorMessage": "Invalid login or password",
		})
		return
	}

	CurrentUser = user

	token, err := utils.GenerateToken(user.Login)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	ctx.SetCookie("Authorization", token, 3600, "/", "", true, false)
	ctx.Set("isAuthenticated", true)
	ctx.Redirect(http.StatusSeeOther, "/home/")

}
