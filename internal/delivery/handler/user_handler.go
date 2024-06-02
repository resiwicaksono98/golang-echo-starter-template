package handler

import (
	"echo-starter-template/internal/adapter/service"
	"echo-starter-template/pkg/utils"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: *service}
}

func (h *UserHandler) FindAll(c echo.Context) error {
	users, err := h.service.FindAll()
	if err != nil {
		return utils.SendErrorResponse(c, err.Error())
	}
	return utils.SendSuccessResponse(c, users, "Berhasil mendapatkan data user")
}
