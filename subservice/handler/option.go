package handler

import (
	"github.com/daifiyum/cat-box/subservice/database"
	"github.com/daifiyum/cat-box/subservice/models"
	"github.com/daifiyum/cat-box/task"

	"github.com/gofiber/fiber/v2"
)

func GetOption(c *fiber.Ctx) error {
	db := database.DB
	options := new(models.Options)
	if err := db.Model(options).Where("name = ?", "options").First(options).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Cannot get option", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Successfully get option", "data": options})
}

func UpdateOption(c *fiber.Ctx) error {
	options := new(models.Options)
	if err := c.BodyParser(options); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Cannot parse body input", "data": nil})
	}
	db := database.DB
	db.Model(options).Where("name = ?", "options").Updates(options)

	// Update delay
	if options.UpdateDelay != "" {
		task.Scheduler(options.UpdateDelay)
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Successfully update option", "data": nil})
}
