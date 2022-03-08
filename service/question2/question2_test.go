package question2

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type ResultTestSuite struct {
	suite.Suite
	service Service
	context echo.Context
	rec     *httptest.ResponseRecorder
	request Request
}

func (suite *ResultTestSuite) SetupTest() {

	suite.service = Service{}

	suite.rec = httptest.NewRecorder()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/cashier", nil)
	suite.context = e.NewContext(req, suite.rec)
	suite.request = Request{
		ProductPrice: 100,
		CustomerPays: 100000,
	}
}

func TestResultTestSuite(t *testing.T) {
	suite.Run(t, new(ResultTestSuite))
}

func (suite *ResultTestSuite) TestNewService() {
	expected := &Service{}
	actual := New()
	suite.Equal(expected, actual, "should be equal")
}

func (suite *ResultTestSuite) Test_Cashier_Status200() {
	_, err := suite.service.Cashier(suite.context, Request{})
	suite.NoError(err)
}

func (suite *ResultTestSuite) Test_Cashier_Coinsstored_Over() {
	_, err := suite.service.Cashier(suite.context, suite.request)
	suite.Error(err)
}
