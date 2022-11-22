package controllers

import (
	"tech-notes-backend/database"
	"tech-notes-backend/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []models.User
	db.Preload(clause.Associations).Find(&users)
	if len(users) == 0 {
		return c.Status(400).JSON(fiber.Map{"message": "No users found"})
	}
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {

	type userBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
		RoleID   uint   `json:"role_id"`
	}

	db := database.DB
	newUser := new(userBody)
	if err := c.BodyParser(newUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})

	}

	hash, err := hashPassword(newUser.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})

	}

	newUser.Password = hash

	user := &models.User{
		Username: newUser.Username,
		Password: newUser.Password,
		RoleID:   newUser.RoleID,
	}

	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "New user " + user.Username + " created"})
}
