package routes

import (
	"github.com/firhan200/taskmanagement/controllers"
	"github.com/firhan200/taskmanagement/middlewares"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(app *fiber.App) {
	tasks := app.Group("/tasks", middlewares.JwtAuthMiddleware)
	tasks.Get("/", controllers.GetTasks)
	tasks.Post("/", controllers.CreateTask)
	// tasks.Get("/:id", controllers.GetTaskById)
	// tasks.Patch("/:id", controllers.UpdateTask)
	// tasks.Delete("/:id", controllers.DeleteTask)
}
