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

	if request.CustomerPays < request.ProductPrice {
		return c.JSON(http.StatusBadRequest, "CustomerPays less ProductPrice")
	}

	res, err := h.service.Cashier(c, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, &res)

}
