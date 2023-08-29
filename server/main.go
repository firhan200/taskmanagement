package main

import (
	"github.com/firhan200/taskmanagement/data"
	"github.com/firhan200/taskmanagement/middlewares"
	"github.com/firhan200/taskmanagement/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	//auto migrate gorm
	data.Migrate()

	//create default for gin
	app := gin.Default()

	app.Use(middlewares.CORSMiddleware)

	//grouping routes based on api version
	v := app.Group("/v1/api")

	//init public routes
	routes.PublicRoutes(v)
	routes.PrivateRoutes(v)

	//run gin server
	app.Run(":8000")
}
