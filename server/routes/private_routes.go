package routes

import (
	"github.com/firhan200/taskmanagement/controllers"
	"github.com/gin-gonic/gin"
)

func PrivateRoutes(app *gin.RouterGroup) {
	profiles := app.Group("profiles")
	profiles.GET("/", controllers.GetProfiles)

	tasks := app.Group("tasks")
	tasks.GET("/", controllers.GetTasks)
	tasks.GET("/:id", controllers.GetTaskById)
}
