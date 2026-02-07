package middleware

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func Recovery(log *zap.Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				log.Error("panic recovered", zap.Any("error", r), zap.String("request_id", c.Locals("request_id").(string)))
			}
		}()
		return c.Next()
	}
}
