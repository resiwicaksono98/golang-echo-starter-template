package main

import (
	"echo-starter-template/internal/adapter/repository"
	"echo-starter-template/internal/adapter/service"
	"echo-starter-template/internal/app/config"
	"echo-starter-template/internal/delivery"
	"echo-starter-template/internal/delivery/handler"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	e := echo.New()

	// Reposptory
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// Router
	delivery.NewRouter(e, userHandler)

	// Start the server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Server.Port)))
}
