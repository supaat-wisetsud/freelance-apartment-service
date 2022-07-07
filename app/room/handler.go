package room

import (
	"fmt"
	"net/http"
	"strconv"

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

	customers, err := h.service.FindAllRoom()

	if err != nil {
		return errorCustomMessage(c, err.Error())
	}

	return c.JSON(http.StatusOK, customers)
}

func (h *Handler) HandlerFindOne(c echo.Context) error {

	sId := c.Param("id")

	id, err := strconv.Atoi(sId)
	if err != nil {
		return errorInvalidBodyRequest(c)
	}

	customer, err := h.service.FindOneRoomByID(uint64(id))

	if err != nil {
		return errorCustomMessage(c, err.Error())
	}

	return c.JSON(http.StatusOK, customer)
}

func (h *Handler) HandlerCreate(c echo.Context) error {

	var req requestUpdate
	if err := c.Bind(&req); err != nil {
		return errorInvalidBodyRequest(c)
	}

	if err := h.service.CreateRoom(req.Name, req.CustomerID, req.Active); err != nil {
		return errorCustomMessage(c, err.Error())
	}

	return c.JSON(http.StatusCreated, response{
		Message: msgCreateRoomSuccess,
	})
}

func (h *Handler) HandlerUpdate(c echo.Context) error {
	sId := c.Param("id")

	id, err := strconv.Atoi(sId)
	if err != nil {
		return errorCustomMessage(c, err.Error())
	}

	var req requestUpdate
	if err := c.Bind(&req); err != nil {
		return errorInvalidBodyRequest(c)
	}

	if err := h.service.UpdateRoom(req.Name, req.CustomerID, req.Active, uint64(id)); err != nil {
		return errorCustomMessage(c, err.Error())
	}

	return c.JSON(http.StatusCreated, response{
		Message: msgUpdateRoomSuccess,
	})
}

func (h *Handler) HandlerRemove(c echo.Context) error {
	sId := c.Param("id")

	id, err := strconv.Atoi(sId)
	if err != nil {
		return errorCustomMessage(c, err.Error())
	}

	if err := h.service.RemoveRoomByID(uint64(id)); err != nil {
		return errorCustomMessage(c, err.Error())
	}

	return c.JSON(http.StatusOK, response{
		Message: msgDeleteRoomSuccess,
	})
}
func (h *Handler) HandlerDestory(c echo.Context) error {
	sId := c.Param("id")

	id, err := strconv.Atoi(sId)
	if err != nil {
		return errorCustomMessage(c, err.Error())
	}

	if err := h.service.DestoryRoomByID(uint64(id)); err != nil {
		return errorCustomMessage(c, err.Error())
	}

	return c.JSON(http.StatusOK, response{
		Message: msgDestoryRoomSuccess,
	})
}

func (h *Handler) HandlerUploadPicture(c echo.Context) error {
	sId := c.Param("id")

	id, err := strconv.Atoi(sId)
	if err != nil {
		return errorCustomMessage(c, err.Error())
	}

	file, err := c.FormFile("picture")
	fmt.Println("file")
	if err != nil {
		return errorCustomMessage(c, err.Error())
	}

	if err := h.service.UpdatePictureByID(file, uint64(id)); err != nil {
		return errorCustomMessage(c, err.Error())
	}

	return c.JSON(http.StatusOK, response{
		Message: msgUpdatePictureSuccess,
	})
}
