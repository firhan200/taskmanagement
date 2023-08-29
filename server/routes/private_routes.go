package routes

import (
	"github.com/firhan200/taskmanagement/controllers"
	"github.com/firhan200/taskmanagement/middlewares"
	"github.com/gin-gonic/gin"
)

func PrivateRoutes(app *gin.RouterGroup) {
	app.Use(middlewares.JwtAuthMiddleware)

	profiles := app.Group("profiles")
	profiles.GET("/", controllers.GetProfiles)

	tasks := app.Group("tasks")
	tasks.GET("/", controllers.GetTasks)
	tasks.POST("/", controllers.CreateTask)
	tasks.GET("/:id", controllers.GetTaskById)
	tasks.PATCH("/:id", controllers.UpdateTask)
	tasks.DELETE("/:id", controllers.DeleteTask)
}
