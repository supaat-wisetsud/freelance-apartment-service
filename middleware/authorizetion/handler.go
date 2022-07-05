package authorizetion

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type handle struct {
	service Service
}

func NewHandler(service Service) handle {
	return handle{service: service}
}

func (h *handle) Handler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header.Get(echo.HeaderAuthorization)
		authToken := strings.TrimSpace(authorization)
		if authToken == "" || strings.ToLower(authToken[:7]) != "bearer " {
			return c.JSON(http.StatusUnauthorized, response{
				Message: msgUnauthorize,
			})
		}
		authToken = strings.TrimSpace(authToken[7:])

		claims, err := h.service.ValidateToken(authToken)

		if err != nil {
			return c.String(http.StatusBadGateway, err.Error())
		}

		t, err := h.service.FindToken(authToken)

		if err != nil {
			return c.String(http.StatusBadGateway, err.Error())
		}
		if !t {
			return c.String(http.StatusUnauthorized, msgUnauthorize)
		}

		c.Set("userID", claims["userID"])
		c.Set("accessTokenKey", authToken)
		return next(c)
	}
}
