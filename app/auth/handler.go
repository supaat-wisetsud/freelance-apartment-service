package auth

import (
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

func (h *Handler) HandlerLogin(c echo.Context) error {

	var req requestLogin

	if err := c.Bind(&req); err != nil {
		// logs.Error(fmt.Sprintf("%s with %v", msgInvalidBodyRequest, err))
		return errorInvalidBodyRequest(c)
	}

	token, err := h.service.generateAccessToken(req.Username, req.Password)

	if err != nil {
		return errorCustomMessage(c, err.Error())
	}

	return c.JSON(http.StatusOK, responseAccessToken{
		IsSuccess:   true,
		AccessToken: token,
	})
}

func (h *Handler) HandlerRegister(c echo.Context) error {

	var req requestRegister

	if err := c.Bind(&req); err != nil {
		return errorInvalidBodyRequest(c)
	}

	if err := h.service.registerUser(req.Username, req.Password, req.Name, req.PhoneNo, req.Email); err != nil {
		return errorCustomMessage(c, err.Error())
	}

	return c.JSON(http.StatusCreated, response{
		Message: msgRegisterSuccess,
	})
}

func (h *Handler) HandlerLogout(c echo.Context) error {
	accessToken := c.Get("accessTokenKey").(string)

	err := h.service.revoke(accessToken)

	if err != nil {
		return errorUnexpected(c)
	}

	return c.JSON(http.StatusOK, response{
		Message: msgRevoke,
	})
}
