package routes

import (
	"github.com/firhan200/taskmanagement/controllers"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(app *gin.RouterGroup) {
	auth := app.Group("auth")
	auth.POST("/login", controllers.Login)
	auth.POST("/register", controllers.Register)
}
