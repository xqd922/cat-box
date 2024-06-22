package handler

import (
	"github.com/daifiyum/cat-box/subservice/database"
	"github.com/daifiyum/cat-box/subservice/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CreateUser new user
func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(models.Users)
	var count int64
	db.Model(user).Count(&count)
	if count > 0 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "error", "message": "User already exists", "data": nil})
	}
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})

	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})

	}

	user.Password = hash
	if err := db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": user})
}

func IsRegistered(c *fiber.Ctx) error {
	type Response struct {
		IsRegistered bool `json:"is_registered"`
	}
	db := database.DB
	user := new(models.Users)
	var count int64
	db.Model(user).Count(&count)
	response := Response{IsRegistered: true}
	if count > 0 {
		return c.JSON(fiber.Map{"status": "success", "message": "User already exists", "data": response})
	}
	response.IsRegistered = false
	return c.JSON(fiber.Map{"status": "success", "message": "User doesn't exist", "data": response})
}
