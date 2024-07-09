package routers

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	SetupHealthRoutes(api)
	SetupTaskRoutes(api)
}
