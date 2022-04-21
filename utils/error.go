package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func BuildErrorResponse(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).SendString(message)
}

func BuildMultipleErrorResponse(c *fiber.Ctx, status int, message string, errs map[string]string) error {
	res := fmt.Errorf(message, errs)
	return c.Status(status).JSON(res)
}
