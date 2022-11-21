package routes

import (
	"tech-notes-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	apiv1 := app.Group("/api/v1")
	apiv1.Get("/", controllers.Healthcheck)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
}
