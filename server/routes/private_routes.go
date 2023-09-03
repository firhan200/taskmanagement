package routes

import (
	"github.com/firhan200/taskmanagement/controllers"
	"github.com/firhan200/taskmanagement/middlewares"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(app *fiber.App) {
	taskHandler := controllers.NewTaskHandler()

	tasks := app.Group("/tasks", middlewares.JwtAuthMiddleware)
	tasks.Get("/", taskHandler.GetTasks())
	tasks.Post("/", taskHandler.CreateTask())
	tasks.Get("/:id", taskHandler.GetTaskById())
	tasks.Patch("/:id", taskHandler.UpdateTask())
	tasks.Delete("/:id", taskHandler.DeleteTask())

	//faker only for test
	faker := app.Group("/generate", middlewares.JwtAuthMiddleware)
	faker.Get("/", taskHandler.GenerateRandomData())
}
