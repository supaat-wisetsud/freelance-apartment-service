package customer

import (
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

	customers, err := h.service.FindAllCustomer()

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

	customer, err := h.service.FindOneCustomerByID(uint64(id))

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

	if err := h.service.CreateCustomer(req.Name, req.CitizenID, req.PhoneNo, req.Email, req.Address); err != nil {
		return errorCustomMessage(c, err.Error())
	}

	return c.JSON(http.StatusCreated, response{
		Message: msgCreateCustomerSuccess,
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

	if err := h.service.UpdateCustomer(req.Name, req.CitizenID, req.PhoneNo, req.Email, req.Address, uint64(id)); err != nil {
		return errorCustomMessage(c, err.Error())
	}

	return c.JSON(http.StatusCreated, response{
		Message: msgCreateCustomerSuccess,
	})
}

func (h *Handler) HandlerRemove(c echo.Context) error {
	sId := c.Param("id")

	id, err := strconv.Atoi(sId)
	if err != nil {
		return errorCustomMessage(c, err.Error())
	}

	if err := h.service.RemoveCustomerByID(uint64(id)); err != nil {
		return errorCustomMessage(c, err.Error())
	}

	return c.JSON(http.StatusOK, response{
		Message: msgDeleteCustomerSuccess,
	})
}
func (h *Handler) HandlerDestory(c echo.Context) error {
	sId := c.Param("id")

	id, err := strconv.Atoi(sId)
	if err != nil {
		return errorCustomMessage(c, err.Error())
	}

	if err := h.service.DestoryCustomerByID(uint64(id)); err != nil {
		return errorCustomMessage(c, err.Error())
	}

	return c.JSON(http.StatusOK, response{
		Message: msgDestoryCustomerSuccess,
	})
}

func (h *Handler) HandlerUploadProfile(c echo.Context) error {
	sId := c.Param("id")

	id, err := strconv.Atoi(sId)
	if err != nil {
		return errorCustomMessage(c, err.Error())
	}

	file, err := c.FormFile("profile")
	if err != nil {
		return err
	}

	if err := h.service.UpdateProfileByID(file, uint64(id)); err != nil {
		return errorCustomMessage(c, err.Error())
	}

	return c.JSON(http.StatusOK, response{
		Message: msgUpdateProfileCustomerSuccess,
	})
}
