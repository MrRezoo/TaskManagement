package routers

import (
	"github.com/MrRezoo/TaskManagement/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupHealthRoutes(router fiber.Router) {
	healthHandler := handlers.NewHealthHandler()

	router.Get("/health/", healthHandler.Check)
}
