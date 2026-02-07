package health

import "github.com/gofiber/fiber/v3"

func Register(r fiber.Router) {
	r.Get("/health", Check)
}
