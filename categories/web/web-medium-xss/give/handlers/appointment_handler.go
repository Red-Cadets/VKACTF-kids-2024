package handlers

import (
	"encoding/base64"
	"example/v3/db"
	"example/v3/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func AppointmentPage(ctx *gin.Context) {

	isAuthenticated := ctx.GetBool("isAuthenticated")
	if isAuthenticated {
		ctx.HTML(http.StatusOK, "appointment.html", nil)
	} else {
		ctx.Redirect(http.StatusTemporaryRedirect, "/login")
	}

}

func AppointmentHandler(ctx *gin.Context) {

	isAuthenticated := ctx.GetBool("isAuthenticated")
	if isAuthenticated {

		date := ctx.PostForm("date")
		time := ctx.PostForm("time")
		doctor := ctx.PostForm("doctor")
		complain := ctx.PostForm("complain")

		db, err := db.OpenDBConnection()
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Error opening connection")
			return
		}

		var existingAppointment models.Appointments
		if err := db.Where("Date = ? AND Time = ? AND Doctor = ?", date, time, doctor).First(&existingAppointment).Error; err == nil {
			ctx.HTML(http.StatusOK, "appointment.html", gin.H{
				"ErrorMessage": "The doctor is busy on the specified date and time",
			})
			return
		}

		user := CurrentUser

		encodedURL := base64.URLEncoding.EncodeToString([]byte(date + time + doctor))

		newAppointment := &models.Appointments{
			Login:      user.Login,
			Date:       date,
			Time:       time,
			Doctor:     doctor,
			Complain:   complain,
			EncodedURL: encodedURL,
		}

		if err := db.Create(&newAppointment).Error; err != nil {
			ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error for add appointment: %v", err))
			return
		}

		ctx.SetCookie("Queue", "we will stand for so long", 3600, "/appointment", "", true, false)
		ctx.Redirect(http.StatusSeeOther, "/appointment/"+encodedURL)
	} else {
		ctx.Redirect(http.StatusTemporaryRedirect, "/login")
	}

}
