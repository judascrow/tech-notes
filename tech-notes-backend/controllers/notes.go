package controllers

import (
	"tech-notes-backend/database"
	"tech-notes-backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllNotes(c *fiber.Ctx) error {
	db := database.DB
	var notes []models.Note
	db.Find(&notes)
	return c.JSON(notes)
}
