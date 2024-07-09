package api

import (
	"github.com/MrRezoo/TaskManagement/api/helpers"
	"github.com/MrRezoo/TaskManagement/api/middlewares"
	"github.com/MrRezoo/TaskManagement/api/routers"
	"github.com/MrRezoo/TaskManagement/config"
	"github.com/MrRezoo/TaskManagement/db"
	"github.com/gofiber/fiber/v2"
	"log"
)

func SetupServer() *fiber.App {
	app := fiber.New()

	app.Use(middlewares.CorsMiddleware())
	app.Use(middlewares.LoggingMiddleware())

	routers.SetupRoutes(app)

	return app
}

func Serve() {
	cfg := config.GetConfig()
	db.ConnectPostgres(&cfg.Postgres)
	app := SetupServer()
	helpers.ListRoutes(app)
	log.Fatal(app.Listen(":" + cfg.Server.Port))

}
