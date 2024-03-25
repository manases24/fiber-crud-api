package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mnsh5/fiber-crud-api/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/hello", handlers.Hello)
}
