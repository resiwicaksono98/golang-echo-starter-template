package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseData struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendSuccessResponse(c echo.Context, data interface{}, message string) error {
	response := ResponseData{
		Success: true,
		Message: message,
		Data:    data,
	}

	return c.JSON(http.StatusOK, response)
}

func SendErrorResponse(c echo.Context, message string) error {
	response := ResponseData{
		Success: false,
		Message: message,
	}

	return c.JSON(http.StatusBadRequest, response)
}
