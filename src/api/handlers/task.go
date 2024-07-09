package handlers

import (
	"github.com/MrRezoo/TaskManagement/api/models"
	"github.com/MrRezoo/TaskManagement/db"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
)

type TaskHandler struct{}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{}
}

func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	task.ID = uuid.New()
	db.DB.Create(&task)
	return c.Status(fiber.StatusCreated).JSON(task)
}

func (h *TaskHandler) GetTasks(c *fiber.Ctx) error {
	var tasks []models.Task

	// Check if DB is initialized
	if db.DB == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database connection not initialized",
		})
	}

	// Fetch tasks from database
	result := db.DB.Find(&tasks)
	if result.Error != nil {
		log.Printf("Failed to fetch tasks: %v", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch tasks",
		})
	}

	return c.JSON(tasks)
}

func (h *TaskHandler) GetTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.Task
	if result := db.DB.First(&task, "id = ?", id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}
	return c.JSON(task)
}

func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.Task
	if result := db.DB.First(&task, "id = ?", id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	db.DB.Save(&task)
	return c.JSON(task)
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if result := db.DB.Delete(&models.Task{}, "id = ?", id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
