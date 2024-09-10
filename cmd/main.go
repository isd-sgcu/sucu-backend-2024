package main

import (
	"github.com/isd-sgcu/sucu-backend-2024/cmd/server"
	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/usecases"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/handlers"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/repositories"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/config"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/database"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/logger"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/s3client"
)

// @title SUCU Backend - API
// @version 0.0.0
// @description  This is an SUCU Backend API in SUCU project.

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	cfg := config.GetConfig()
	db := database.NewGormDatabase(cfg)
	s3 := s3client.NewS3Client(cfg)
	logger := logger.NewLogger(cfg)

	repositories := repositories.NewRepository(cfg, db, s3)
	usecases := usecases.NewService(repositories, cfg, logger)
	handlers := handlers.NewHandler(usecases)

	servers := server.NewFiberHttpServer(cfg, logger, handlers)

	servers.Start()
}
