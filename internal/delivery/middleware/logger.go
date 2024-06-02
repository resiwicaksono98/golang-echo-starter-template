package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoggerConfig() middleware.LoggerConfig {
	return middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}
}

func InitLogger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(LoggerConfig())
}
