package handlers

import (
	"github.com/MrRezoo/TaskManagement/api/helpers"
	"github.com/MrRezoo/TaskManagement/api/models"
	"github.com/MrRezoo/TaskManagement/api/validations"
	"github.com/MrRezoo/TaskManagement/db"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TaskHandler struct{}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{}
}

func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.GenerateResponseWithError(nil, false, fiber.StatusBadRequest, err))
	}

	// Validate the task
	if err := validations.Validate.Struct(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.GenerateResponseWithValidationErrors(nil, false, fiber.StatusBadRequest, err))
	}

	task.ID = uuid.New()
	if result := db.DB.Create(&task); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(helpers.GenerateResponseWithError(nil, false, fiber.StatusInternalServerError, result.Error))
	}
	return c.Status(fiber.StatusCreated).JSON(helpers.GenerateBaseResponse(task, true, fiber.StatusCreated))
}

func (h *TaskHandler) GetTasks(c *fiber.Ctx) error {
	var tasks []models.Task
	if result := db.DB.Find(&tasks); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(helpers.GenerateResponseWithError(nil, false, fiber.StatusInternalServerError, result.Error))
	}
	return c.JSON(helpers.GenerateBaseResponse(tasks, true, fiber.StatusOK))
}

func (h *TaskHandler) GetTask(c *fiber.Ctx) error {
	id := c.Params("id")

	if _, err := uuid.Parse(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.GenerateResponseWithValidationErrors(nil, false, fiber.StatusBadRequest, err))
	}

	var task models.Task
	if result := db.DB.First(&task, "id = ?", id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(helpers.GenerateResponseWithError(nil, false, fiber.StatusNotFound, result.Error))
	}
	return c.JSON(helpers.GenerateBaseResponse(task, true, fiber.StatusOK))
}

func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")

	if _, err := uuid.Parse(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.GenerateResponseWithValidationErrors(nil, false, fiber.StatusBadRequest, err))
	}

	var task models.Task
	if result := db.DB.First(&task, "id = ?", id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(helpers.GenerateResponseWithError(nil, false, fiber.StatusNotFound, result.Error))
	}

	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.GenerateResponseWithError(nil, false, fiber.StatusBadRequest, err))
	}

	if err := validations.Validate.Struct(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.GenerateResponseWithValidationErrors(nil, false, fiber.StatusBadRequest, err))
	}

	if result := db.DB.Save(&task); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(helpers.GenerateResponseWithError(nil, false, fiber.StatusInternalServerError, result.Error))
	}
	return c.JSON(helpers.GenerateBaseResponse(task, true, fiber.StatusOK))
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")

	if _, err := uuid.Parse(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.GenerateResponseWithValidationErrors(nil, false, fiber.StatusBadRequest, err))
	}

	if result := db.DB.Delete(&models.Task{}, "id = ?", id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(helpers.GenerateResponseWithError(nil, false, fiber.StatusNotFound, result.Error))
	}
	return c.Status(fiber.StatusNoContent).JSON(helpers.GenerateBaseResponse(nil, true, fiber.StatusNoContent))
}
