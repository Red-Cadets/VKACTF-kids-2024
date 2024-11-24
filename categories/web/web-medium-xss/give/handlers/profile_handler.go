package handlers

import (
	"example/v3/db"
	"example/v3/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProfilePage(ctx *gin.Context) {

	isAuthenticated := ctx.GetBool("isAuthenticated")
	if isAuthenticated {

		user := CurrentUser
		if user == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		db, err := db.OpenDBConnection()
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Error opening connection")
			return
		}

		var appointments []models.Appointments
		if err := db.Where("Login = ?", user.Login).Find(&appointments).Error; err != nil {
			ctx.String(http.StatusInternalServerError, "Error fetching appointments: %v", err)
			return
		}

		ctx.HTML(http.StatusOK, "profile.html", gin.H{
			"Name":         user.Name,
			"Surname":      user.Surname,
			"DateOfBirth":  user.DateOfBirthday,
			"Email":        user.Email,
			"PhoneNumber":  user.PhoneNumber,
			"Passport":     user.Passport,
			"Login":        user.Login,
			"AvatarURL":    "https://ui-avatars.com/api/?name=" + user.Name + "+" + user.Surname + "&background=random",
			"Appointments": appointments,
		})
	} else {
		ctx.Redirect(http.StatusTemporaryRedirect, "/login")
	}

}
