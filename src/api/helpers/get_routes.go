package helpers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func ListRoutes(app *fiber.App) {
	fmt.Println("List of registered routes:")
	routes := app.GetRoutes()
	for _, route := range routes {

		fmt.Printf("Request: %s -> %s \n", route.Method, route.Path)
	}
}
