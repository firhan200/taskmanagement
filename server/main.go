package main

import (
	"github.com/firhan200/taskmanagement/data"
	"github.com/firhan200/taskmanagement/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	//auto migrate schema to db
	data.Migrate()

	app := fiber.New()

	//setup cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://localhost:5173, http://localhost:5173",
		AllowHeaders:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	routes.PrivateRoutes(app)
	routes.PublicRoutes(app)

	app.Listen(":8000")
}
