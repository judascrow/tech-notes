package controllers

import (
	"tech-notes-backend/database"
	"tech-notes-backend/models"

	"github.com/gofiber/fiber/v2"
)

func Healthcheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "API is Online!"})
}

func GetAllRoles(c *fiber.Ctx) error {
	db := database.DB
	var roles []models.Role
	db.Find(&roles)
	return c.JSON(roles)
}
