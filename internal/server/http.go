package server

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/priyanshu334/go-template/internal/config"
	"github.com/priyanshu334/go-template/internal/middleware"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewHTTPServer(cfg *config.Config, log *zap.Logger, db *gorm.DB) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "go template",
	})
	app.Use(middleware.RequestID())
	app.Use(middleware.Recovery(log))

	RegisterRoutes(app, db)

	log.Info("fiber server intilization")
	return app
}

func Start(app *fiber.App, cfg *config.Config, log *zap.Logger) {
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Info("http server starting", zap.String("addr", addr))
	if err := app.Listen(addr); err != nil {
		log.Fatal("failed to start http server", zap.Error(err))
	}
}
