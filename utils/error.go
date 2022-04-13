package utils

import "github.com/gofiber/fiber/v2"

func BuildErrorResponse(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).SendString(message)
}
