package handlers

import (
	"example/v3/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()
	router.Use(middleware.AuthMiddleware())

	router.GET("/home/", HomePage)
	router.GET("/login/", middleware.RedirectIfAuthenticated(), LoginPage)
	router.POST("/login/", middleware.RedirectIfAuthenticated(), LoginHandler)
	router.GET("/signup/", middleware.RedirectIfAuthenticated(), SignUpPage)
	router.POST("/signup/", middleware.RedirectIfAuthenticated(), SignUpHandler)

	protected := router.Group("/")
	{
		protected.GET("/profile/", ProfilePage)
		protected.GET("/logout/", LogoutPage)
		protected.GET("/appointment/", AppointmentPage)
		protected.POST("/appointment/", AppointmentHandler)
		protected.GET("/appointment/:encodedURL", AppointmentURLHandler)
	}

	return router
}
