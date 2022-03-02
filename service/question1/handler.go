package question1

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

func (h *Handler) CalculateDataset(c echo.Context) error {

	res := h.service.CalculateDataset(c)

	return c.JSON(http.StatusOK, &res)

}
