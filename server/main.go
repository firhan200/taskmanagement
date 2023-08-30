package main

import (
	"github.com/firhan200/taskmanagement/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// func main() {
// 	//auto migrate gorm
// 	data.Migrate()

// 	//create default for gin
// 	app := gin.Default()

// 	app.Use(cors.Default())
// 	//grouping routes based on api version
// 	v := app.Group("/v1/api")

// 	//init public routes
// 	routes.PrivateRoutes(v)
// 	routes.PublicRoutes(v)

// 	//run gin server
// 	app.Run(":8000")
// }

func main() {
	app := fiber.New()

	// Or extend your config for customization
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
