package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type HealthController struct {
}

func (c HealthController) HandleHealthCheck() fiber.Handler {
	return func(context *fiber.Ctx) error {
		result := &fiber.Map{
			"result": "OK",
		}

		return context.Status(fiber.StatusOK).JSON(result)
	}
}
