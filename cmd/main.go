package main

import (
	"fmt"

	"github.com/isd-sgcu/sucu-backend-2024/cmd/server"
	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/usecases"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/handlers"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/repositories"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/config"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/database"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/logger"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/s3client"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/validator"
)

// @title SUCU Backend - API
// @version 0.1.0
// @description  This is an SUCU Backend API in SUCU project.

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and the token
func main() {
	cfg := config.GetConfig()
	db := database.NewGormDatabase(cfg)
	s3 := s3client.NewS3Client(cfg)
	logger := logger.NewLogger(cfg)
	validator, err := validator.NewDtoValidator()
	if err != nil {
		panic(fmt.Sprintf("Failed to create dto validator: %v", err))
	}

	repositories := repositories.NewRepository(cfg, db, s3)
	usecases := usecases.NewUsecase(repositories, cfg, logger)
	handlers := handlers.NewHandler(usecases, validator)

	servers := server.NewFiberHttpServer(cfg, logger, handlers)

	servers.Start()
}
