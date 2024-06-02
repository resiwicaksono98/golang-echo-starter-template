package router

import (
	"echo-starter-template/internal/delivery/handler"
	"echo-starter-template/internal/delivery/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, userhandler *handler.UserHandler) {
	e.Use(middleware.InitLogger())
	e.Use(middleware.InitCors())
	e.HTTPErrorHandler = middleware.ErrorHandler

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users", userhandler.FindAll)
}
