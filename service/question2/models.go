package question2

type Request struct {
	ProductPrice float64 `json:"productPrice"`
	CustomerPays float64 `json:"customerPays"`
}

type Coinsstoreds struct {
	Coinsstored []Coinsstored `json:"coinsstored"`
}

type Coinsstored struct {
	Description string  `json:"description"`
	Value       float64 `json:"value"`
	Amount      int     `json:"amount"`
	Type        string  `json:"type"`
}

type Response struct {
	CalculateTheChangeMoney float64      `json:"calculateTheChangeMoney"`
	Coinsstoreds            Coinsstoreds `json:"coinsstoreds"`
}
