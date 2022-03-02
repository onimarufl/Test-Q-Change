package question2

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type handlerTestSuite struct {
	suite.Suite
	handler Handler
	service Servicer
	rec     *httptest.ResponseRecorder
	context echo.Context
	request Request
}

func (suite *handlerTestSuite) SetupTest() {
	suite.service = new(servicerMock)

	suite.handler = Handler{
		service: suite.service,
	}
	suite.rec = httptest.NewRecorder()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/cashier", nil)
	suite.context = e.NewContext(req, suite.rec)

}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(handlerTestSuite))
}

func (suite *handlerTestSuite) TestNewHandler() {
	expected := &suite.handler
	actual := NewHandler(suite.service)
	suite.Equal(expected, actual, "they should be equal")
}

func (suite *handlerTestSuite) TestHandler_Cashier_Status200() {
	suite.service.(*servicerMock).On("Cashier", suite.context, suite.request).Return(&Response{}, nil)
	if suite.NoError(suite.handler.Cashier(suite.context)) {
		suite.Equal(http.StatusOK, suite.context.Response().Status)
	}

}

func (suite *handlerTestSuite) TestHandler_Cashier_Status500() {
	suite.service.(*servicerMock).On("Cashier", suite.context, suite.request).Return(&Response{}, errors.New("something"))
	if suite.NoError(suite.handler.Cashier(suite.context)) {
		suite.Equal(http.StatusInternalServerError, suite.context.Response().Status)
	}
}
