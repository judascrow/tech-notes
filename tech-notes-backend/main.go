package main

import (
	"log"
	"tech-notes-backend/configs"
	"tech-notes-backend/database"
	"tech-notes-backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	port := configs.Config("SERVER_PORT")

	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${status} | ${latency} | ${method} | ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
	}))
	app.Use(cors.New())
	app.Use(recover.New())

	database.ConnectDB()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":" + port))
}
