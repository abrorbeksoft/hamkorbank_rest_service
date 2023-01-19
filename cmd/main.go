package main

import (
	"github.com/gin-gonic/gin"
	"rest_service/api"
	"rest_service/api/handlers"
	"rest_service/config"
	"rest_service/pkg/logger"
)

func main() {
	cfg := config.Load()

	var loggerLevel = new(string)
	*loggerLevel = logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		*loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.NewLogger("epa_go_api_gateway", *loggerLevel)
	defer func() {
		err := logger.Cleanup(log)
		if err != nil {
			return
		}
	}()

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	h := handlers.NewHandler(cfg, log)

	api.SetUpAPI(r, h, cfg)

	if err := r.Run(cfg.HTTPPort); err != nil {
		return
	}
}
