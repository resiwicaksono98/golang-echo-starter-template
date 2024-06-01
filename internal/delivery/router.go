package delivery

import (
	"echo-starter-template/internal/delivery/handler"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, userhandler *handler.UserHandler) {
	e.GET("/users", userhandler.FindAll)
}
