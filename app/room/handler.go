package room

import "github.com/labstack/echo/v4"

type Handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return Handler{
		service: service,
	}
}

func (h *Handler) HandlerFindAll(c echo.Context) error {

	return c.String(200, "")
}

func (h *Handler) HandlerFindOne(c echo.Context) error {

	return c.String(200, "")
}

func (h *Handler) HandlerCreate(c echo.Context) error {
	return c.String(200, "")
}

func (h *Handler) HandlerUpdate(c echo.Context) error {
	return c.String(200, "")
}

func (h *Handler) HandlerRemove(c echo.Context) error {
	return c.String(200, "")
}
func (h *Handler) HandlerDestory(c echo.Context) error {
	return c.String(200, "")
}

func (h *Handler) HandlerUploadProfile(c echo.Context) error {
	return c.String(200, "")
}
