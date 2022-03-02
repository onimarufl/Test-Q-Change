package question2

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Servicer
}

func NewHandler(service Servicer) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Cashier(c echo.Context) error {

	request := Request{}
	if err := c.Bind(&request); err != nil {
		return echo.ErrBadRequest
	}

	res, err := h.service.Cashier(c, request)
	if err != nil {
		return echo.NewHTTPError(echo.ErrInternalServerError.Code, err.Error())
	}
	return c.JSON(http.StatusOK, &res)

}
