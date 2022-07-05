package room

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func errorUnexpected(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, response{Message: msgErrorUnexpected})
}

func errorInvalidBodyRequest(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, response{
		Message: msgInvalidBodyRequest,
	})
}

func errorCustomMessage(c echo.Context, message string) error {
	return c.JSON(http.StatusBadRequest, response{
		Message: message,
	})
}
