package health

import "github.com/gofiber/fiber/v3"

func Check(c fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status": "ok",
	})
}
