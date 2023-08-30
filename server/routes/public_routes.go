package routes

import (
	"github.com/firhan200/taskmanagement/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	auth := app.Group("auth")
	auth.Post("/login", controllers.Login)
	auth.Post("/register", controllers.Register)
}
