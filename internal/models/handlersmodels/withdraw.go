package handlersmodels

import (
	"context"
	"errors"
	"github.com/shopspring/decimal"
)

var ErrNegativeBalanceValue = errors.New("Insufficient funds to debit bonuses")

var ErrAddHistoryBalance = errors.New("error when adding data to the table history_balance")

var ErrDuplicateOrderNumberHistoryBalance = errors.New("this order number is already in the table history_balance")

var ErrDuplicateOrderNumberOrders = errors.New("this order number is already in the table orders")

type ReqWithdraw struct {
	OrderNumber string          `json:"order"`
	SumWithdraw decimal.Decimal `json:"sum"`
	Ctx         context.Context
}
