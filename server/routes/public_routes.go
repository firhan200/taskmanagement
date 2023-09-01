package routes

import (
	"github.com/firhan200/taskmanagement/controllers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PublicRoutes(app *fiber.App, db *gorm.DB) {
	loginHandler := controllers.NewLoginHandler(db)

	auth := app.Group("auth")
	auth.Post("/login", loginHandler.Login())
	auth.Post("/register", loginHandler.Register())
}
