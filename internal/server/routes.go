package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/priyanshu334/go-template/internal/module/health"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")

	health.Register(api)
}
