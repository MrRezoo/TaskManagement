package routers

import (
	"github.com/MrRezoo/TaskManagement/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupTaskRoutes(router fiber.Router) {
	taskHandler := handlers.NewTaskHandler()

	router.Post("/tasks", taskHandler.CreateTask)
	router.Get("/tasks", taskHandler.GetTasks)
	router.Get("/tasks/:id", taskHandler.GetTask)
	router.Patch("/tasks/:id", taskHandler.UpdateTask)
	router.Delete("/tasks/:id", taskHandler.DeleteTask)
}
