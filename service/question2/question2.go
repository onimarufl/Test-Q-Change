package question2

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

type Servicer interface {
	Cashier(c echo.Context, request Request) (*Response, error)
}

func (s *Service) Cashier(c echo.Context, request Request) (*Response, error) {
	response := Response{}
	coinsstoreds := Coinsstoreds{}
	file, _ := ioutil.ReadFile("./mockFile/coinsstored.json")
	_ = json.Unmarshal([]byte(file), &coinsstoreds)

	calculateTheChangeMoney := request.CustomerPays - request.ProductPrice
	calculateTheChangeMoneyResponse := calculateTheChangeMoney

	sumTotalValue := sumTotalValue(coinsstoreds)
	if sumTotalValue < calculateTheChangeMoney {
		return nil, errors.New("Coinsstored Over")
	}

	coinsstoredsResponse := Coinsstoreds{}

	for _, v := range coinsstoreds.Coinsstored {
		if v.Value <= calculateTheChangeMoney && calculateTheChangeMoney > 0 {
			coinsstoredResponse := Coinsstored{}
			coinsstoredResponse.Description = v.Description
			coinsstoredResponse.Value = v.Value
			coinsstoredResponse.Type = v.Type
			for i := v.Value; i <= calculateTheChangeMoney && calculateTheChangeMoney > 0 && coinsstoredResponse.Amount < v.Amount; {
				coinsstoredResponse.Amount++
				calculateTheChangeMoney = calculateTheChangeMoney - v.Value
			}
			coinsstoredsResponse.Coinsstored = append(coinsstoredsResponse.Coinsstored, coinsstoredResponse)
		}

	}

	response.CalculateTheChangeMoney = calculateTheChangeMoneyResponse
	response.Coinsstoreds = coinsstoredsResponse

	return &response, nil

}

func sumTotalValue(c Coinsstoreds) float64 {
	var totalValue float64
	for _, v := range c.Coinsstored {
		totalValue = totalValue + (v.Value * float64(v.Amount))
	}
	return totalValue
}
