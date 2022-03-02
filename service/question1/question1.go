package question1

import "github.com/labstack/echo/v4"

type Service struct {
}

func New() *Service {
	return &Service{}
}

type Servicer interface {
	CalculateDataset(c echo.Context) string
}

func (s *Service) GetTestService(c echo.Context) string {

	return "test"

}
