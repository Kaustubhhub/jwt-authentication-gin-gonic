package routes

import (
	"jwt_authentication_gin_gonic/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/signup", controllers.Signup)
	incomingRoutes.POST("/signin", controllers.Signin)
}
