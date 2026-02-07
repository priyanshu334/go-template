package app

import (
	"github.com/priyanshu334/go-template/internal/config"
	"github.com/priyanshu334/go-template/internal/database"
	"github.com/priyanshu334/go-template/internal/logger"
	"github.com/priyanshu334/go-template/internal/server"
	"github.com/priyanshu334/go-template/internal/shutdown"
	"go.uber.org/zap"
)

func Start() {
	cfg := config.Load()

	log := logger.New(cfg.Env)
	defer log.Sync()
	db := database.Connect(cfg.Database)
	database.AutoMigrate(db)
	httpApp := server.NewHTTPServer(cfg, log, db)
	go shutdown.Graceful(httpApp, log)
	server.Start(httpApp, cfg, log)

	log.Info("application starting", zap.String("env", cfg.Env), zap.Int("port", cfg.Server.Port))
}
