package routes

import (
	"jwt_authentication_gin_gonic/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.Getusers)
	incomingRoutes.GET("/users/:userId", controllers.Getuser)
}
