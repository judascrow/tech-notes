package main

import (
	"log"
	"tech-notes-backend/configs"
	"tech-notes-backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	port := configs.Config("SERVER_PORT")

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":" + port))
}
