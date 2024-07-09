package handlers

import "github.com/gofiber/fiber/v2"

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler { return &HealthHandler{} }

func (h *HealthHandler) Check(context *fiber.Ctx) error {
	return context.JSON(fiber.Map{"message": "Boom Boom 💥", "success": true})
}
