package routes

import (
	"tech-notes-backend/controllers"
	"tech-notes-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	apiv1 := app.Group("/api/v1")
	apiv1.Get("/", controllers.Healthcheck)

	auth := apiv1.Group("/auth")
	auth.Post("", controllers.Login)

	roles := apiv1.Group("/roles")
	roles.Get("", controllers.GetAllRoles)

	users := apiv1.Group("/users")
	users.Get("", middlewares.Protected(), controllers.GetAllUsers)
	users.Post("", controllers.CreateUser)

	notes := apiv1.Group("/notes")
	notes.Get("", controllers.GetAllNotes)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
}
