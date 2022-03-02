package question2

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type servicerMock struct {
	mock.Mock
}

func (s *servicerMock) Cashier(c echo.Context, request Request) (*Response, error) {
	args := s.Called(c, request)
	return args.Get(0).(*Response), args.Error(1)

}
