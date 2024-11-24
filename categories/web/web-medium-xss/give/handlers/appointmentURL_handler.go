package handlers

import (
	"example/v3/db"
	"example/v3/models"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AppointmentURLHandler(ctx *gin.Context) {

	encodedURL := ctx.Param("encodedURL")

	db, err := db.OpenDBConnection()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Ошибка подключения к базе данных")
		return
	}

	var appointment *models.Appointments
	if err := db.Where("EncodedURL = ?", encodedURL).First(&appointment).Error; err != nil {
		ctx.String(http.StatusNotFound, "Заявка не найдена")
		return
	}

	ctx.HTML(http.StatusOK, "appointmentURL.html", gin.H{
		"Date":     template.HTML(appointment.Date),
		"Time":     template.HTML(appointment.Time),
		"Doctor":   template.HTML(appointment.Doctor),
		"Complain": template.HTML(appointment.Complain),
	})
}
