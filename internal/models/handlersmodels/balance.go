package handlersmodels

import "github.com/shopspring/decimal"

type RespUserBalance struct {
	UserBalance      decimal.Decimal `json:"current"`
	WithdrawnBalance decimal.Decimal `json:"withdrawn"`
}
