package handler

import (
	"github.com/daifiyum/cat-box/singbox"

	"github.com/gofiber/fiber/v2"
)

// StopSingbox query all subscribe
func StopSingbox(c *fiber.Ctx) error {
	err := singbox.Stop()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "can't stop singbox", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "success stop singbox", "data": nil})
}
