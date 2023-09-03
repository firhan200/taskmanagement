package routes

import (
	"github.com/firhan200/taskmanagement/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	loginHandler := controllers.NewLoginHandler()

	auth := app.Group("auth")
	auth.Post("/login", loginHandler.Login())
	auth.Post("/register", loginHandler.Register())
}
