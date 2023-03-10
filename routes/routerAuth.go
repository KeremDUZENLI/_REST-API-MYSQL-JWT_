package routes

import (
	"jwt-project/controller"

	"github.com/gin-gonic/gin"
)

func AuthenticationRoutes(routes *gin.Engine) {
	routes.POST("/person/signup", controller.SignUp)
	routes.POST("/person/login", controller.Login)
}
