package routes

import (
	"jwt_authentication_gin_gonic/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/signup", controllers.Signup)
	incomingRoutes.GET("/signin", controllers.Signin)
}