package logs

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return Handler{
		service: service,
	}
}

func (h *Handler) HandlerFindAll(c echo.Context) error {

	result, err := h.service.FindAllLogs()
	if err != nil {
		log.Println(err.Error())
		return errorUnexpected(c)
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *Handler) HandlerCreate(c echo.Context) error {
	var req requestUpdate
	if err := c.Bind(&req); err != nil {
		return errorInvalidBodyRequest(c)
	}

	if err := h.service.CreateLogs(*req.CustomerID, *req.RoomID); err != nil {
		return errorCustomMessage(c, err.Error())
	}

	return c.JSON(http.StatusCreated, response{
		Message: msgCreateLogSuccess,
	})
}
