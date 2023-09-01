package main

import (
	"os"

	"github.com/firhan200/taskmanagement/data"
	"github.com/firhan200/taskmanagement/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	//get new connection, call this once to create only 1 connection pool
	db := data.NewConnection()
	//auto migrate schema to db
	data.Migrate(db)

	app := fiber.New()

	origins := os.Getenv("ALLOWED_ORIGIN")

	//setup cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowHeaders:     "Authorization, Content-Type, Accept",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	routes.PrivateRoutes(app, db)
	routes.PublicRoutes(app, db)

	app.Listen(":8000")
}
