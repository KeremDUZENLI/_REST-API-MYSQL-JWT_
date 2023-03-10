package routes

import (
	"jwt-project/controller"

	"jwt-project/middleware"

	"github.com/gin-gonic/gin"
)

func PersonRoutes(routes *gin.Engine) {
	routes.Use(middleware.Autheticate())
	routes.GET("/person/:userid", controller.GetUser)
	routes.GET("/personall", controller.GetUsers)
}
