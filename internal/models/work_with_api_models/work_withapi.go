package work_with_api_models

import "github.com/shopspring/decimal"

type ReqAPIOrders struct {
	OrderNumber string  `json:"order"`
	Goods       []Goods `json:"goods"`
}

type Goods struct {
	Description string          `json:"description"`
	Price       decimal.Decimal `json:"price"`
}
