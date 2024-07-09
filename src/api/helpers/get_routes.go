package helpers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"sort"
)

func ListRoutes(app *fiber.App) {
	log.Println("List of registered routes:")
	routes := app.GetRoutes()
	pathMethods := make(map[string][]string)
	allowedMethods := []string{"GET", "POST", "PUT", "PATCH", "DELETE"}

	for _, route := range routes {
		for _, allowedMethod := range allowedMethods {
			if route.Method == allowedMethod {
				pathMethods[route.Path] = append(pathMethods[route.Path], route.Method)
				break
			}
		}
	}

	for path, methods := range pathMethods {
		sort.Strings(methods)
		log.Printf("%s -> %s\n", path, methods)
	}
}
