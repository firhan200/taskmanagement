package routes

import (
	"github.com/firhan200/taskmanagement/controllers"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(app *fiber.App) {
	//app.Use(middlewares.JwtAuthMiddleware)

	tasks := app.Group("/v1/api/tasks")
	tasks.Get("/", controllers.GetTasks)
	// tasks.Post("/", controllers.CreateTask)
	// tasks.Get("/:id", controllers.GetTaskById)
	// tasks.Patch("/:id", controllers.UpdateTask)
	// tasks.Delete("/:id", controllers.DeleteTask)
}
