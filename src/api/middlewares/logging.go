package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

const (
	colorReset       = "\033[0m"
	colorCyan        = "\033[36m"
	colorRed         = "\033[31m"
	bgColorLightBlue = "\033[44m"
)

func LoggingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()

		statusCode := c.Response().StatusCode()
		color := colorCyan
		if statusCode >= 400 && statusCode < 500 {
			color = colorRed
		}

		log.Printf("%s%s %s %d %s%s%s", color, c.Method(), c.Path(), statusCode, bgColorLightBlue, colorReset, colorReset)

		return err
	}
}
